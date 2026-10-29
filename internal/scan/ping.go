package scan

import (
	"context"
	"net"
	"os"
	"sync"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

// Sweep sends an ICMP Echo Request to every address in ips and returns the
// round-trip time of each one that replies, keyed by IP. It waits up to
// timeout for replies, or until ctx is cancelled, and returns whatever replies
// arrived by then.
//
// One socket is shared by everything: a single goroutine reads replies while
// the calling goroutine sends the requests.
func Sweep(ctx context.Context, ips []net.IP, timeout time.Duration) (map[string]time.Duration, error) {
	// "udp4" is Go's name for an unprivileged ICMP socket, which works without
	// sudo on macOS. The packets are real ICMP, not UDP, but addresses on this
	// socket are *net.UDPAddr.
	conn, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// sent is written by the send loop and read by the receiver goroutine, so
	// both maps are only touched while holding mu.
	var mu sync.Mutex
	sent := make(map[string]time.Time)
	results := make(map[string]time.Duration)

	// One deadline for the whole sweep. ReadFrom ignores ctx, so on Ctrl+C the
	// deadline is moved to now, which unblocks the receiver straight away.
	conn.SetReadDeadline(time.Now().Add(timeout))
	stop := context.AfterFunc(ctx, func() {
		conn.SetReadDeadline(time.Now())
	})
	defer stop()

	done := make(chan struct{})
	go func() {
		defer close(done)
		buf := make([]byte, 1500)
		for {
			n, peer, err := conn.ReadFrom(buf)
			if err != nil {
				return
			}
			// 1 is the protocol number for ICMP over IPv4. Other ICMP traffic
			// (unreachable messages, other programs' pings) can arrive on this
			// socket too, so anything that isn't an Echo Reply is skipped.
			rm, err := icmp.ParseMessage(1, buf[:n])
			if err != nil || rm.Type != ipv4.ICMPTypeEchoReply {
				continue
			}
			ip := peer.(*net.UDPAddr).IP.String()
			mu.Lock()
			if start, ok := sent[ip]; ok {
				if _, seen := results[ip]; !seen {
					results[ip] = time.Since(start)
				}
			}
			mu.Unlock()
		}
	}()

	// WriteTo doesn't wait for a reply, so sending to every host takes about a
	// millisecond and needs no extra goroutines.
	for i, ip := range ips {
		if ctx.Err() != nil {
			break
		}
		// ID identifies this program and Seq numbers each request. Marshal
		// fills in the checksum.
		msg := icmp.Message{
			Type: ipv4.ICMPTypeEcho,
			Body: &icmp.Echo{
				ID:   os.Getpid() & 0xffff,
				Seq:  i,
				Data: []byte("gonetsuite"),
			},
		}
		b, err := msg.Marshal(nil)
		if err != nil {
			return nil, err
		}
		mu.Lock()
		sent[ip.String()] = time.Now()
		mu.Unlock()
		conn.WriteTo(b, &net.UDPAddr{IP: ip})
	}
	<-done
	return results, nil
}
