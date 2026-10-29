package command

import (
	"context"
	"io"
	"net"
)

// ScanCommand will scan you local subnet for other devices
func ScanCommand(ctx context.Context, args []string, in io.Reader, out io.Writer) error {
	//implement ping sweep using ICMP echo requests, then later on for a ARP sweep maybe when build sniff?

	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	var localIp *net.IPNet

outer:
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			if ipNet.IP.IsUnspecified() {
				continue
			}
			ip4 := ipNet.IP.To4()
			if ip4 == nil {
				continue
			}
			localIp = ipNet
			break outer
		}
	}

	mask := localIp.Mask
	network := localIp.IP.Mask(mask)

	return nil
}
