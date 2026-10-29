package command

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"slices"
	"time"

	"github.com/shabebe11/GoNetSuite/internal/scan"
)

// ScanCommand will scan you local subnet for other devices
func ScanCommand(ctx context.Context, args []string, in io.Reader, out io.Writer) error {
	//implement ping sweep using ICMP echo requests, then later on for a ARP sweep maybe when build sniff?

	localIp, err := scan.LocalSubnet()
	if err != nil {
		return err
	}

	ips, err := scan.Hosts(localIp)
	if err != nil {
		return err
	}

	fmt.Fprintf(out, "> Scanning %s (%d hosts)...\n", localIp, len(ips))

	results, err := scan.Sweep(ctx, ips, 3*time.Second)
	if err != nil && ctx.Err() == nil {
		return fmt.Errorf("error scanning %s: %w", localIp, err)
	}

	hosts := make([]net.IP, 0, len(results))
	for ip := range results {
		hosts = append(hosts, net.ParseIP(ip).To4())
	}
	slices.SortFunc(hosts, func(a, b net.IP) int {
		return bytes.Compare(a, b)
	})

	for _, ip := range hosts {
		fmt.Fprintf(out, "  %-15s  %v\n", ip, results[ip.String()].Round(100*time.Microsecond))
	}
	fmt.Fprintf(out, "> %d hosts up\n", len(hosts))
	return nil
}
