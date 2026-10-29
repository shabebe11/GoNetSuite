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

func Sweep(ctx context.Context, ips []net.IP, timeout time.Duration) (map[string]time.Duration, error) {
	conn, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var mu sync.Mutex
	sent := make(map[string]time.Time)
	results := make(map[string]time.Duration)

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

	for i, ip := range ips {
		if ctx.Err() != nil {
			break
		}
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
