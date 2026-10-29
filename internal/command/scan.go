package command

import (
	"context"
	"fmt"
	"io"
	"net"
)

// ScanCommand will scan you local subnet for other devices
func ScanCommand(ctx context.Context, args []string, in io.Reader, out io.Writer) error {
	//implement ping sweep using ICMP echo requests, then later on for a ARP sweep maybe when build sniff?

	localIp, err := localSubnet()
	if err != nil {
		return err
	}

	mask := localIp.Mask
	network := localIp.IP.Mask(mask)

	return nil
}

func localSubnet() (*net.IPNet, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok || ipNet.IP.IsUnspecified() || ipNet.IP.To4() == nil {
				continue
			}
			return ipNet, nil
		}
	}
	return nil, fmt.Errorf("no suitable local IP address found")
}
