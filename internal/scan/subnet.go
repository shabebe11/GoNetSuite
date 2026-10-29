package scan

import (
	"encoding/binary"
	"fmt"
	"net"
)

func LocalSubnet() (*net.IPNet, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
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

func Hosts(ipNet *net.IPNet) ([]net.IP, error) {
	ip4 := ipNet.IP.To4()
	ones, bits := ipNet.Mask.Size()
	if bits-ones > 10 {
		return nil, fmt.Errorf("subnet /%d too large to sweep", ones)
	}

	network := binary.BigEndian.Uint32(ip4.Mask(ipNet.Mask))
	size := uint32(1) << (bits - ones)

	var ips []net.IP
	for i := uint32(1); i < size-1; i++ {
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, network+i)
		if ip.Equal(ip4) {
			continue
		}
		ips = append(ips, ip)
	}
	return ips, nil
}
