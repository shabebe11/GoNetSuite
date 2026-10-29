package scan

import (
	"context"
	"net"
	"sync"
	"time"

	"golang.org/x/net/icmp"
)

func Sweep(ctx context.Context, ips []net.IP, timeout time.Duration) (map[string]time.Duration, error) {
	conn, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		return nil, err
	}

	var mu sync.Mutex
	sent := make(map[string]time.Duration)
	results := make(map[string]time.Duration)

	conn.SetReadDeadline(time.Now().Add(timeout))
	stop := context.AfterFunc(ctx, func() {
		conn.SetReadDeadline(time.Now())
	})
	defer stop()

	defer conn.Close()
	return results, nil
}
