package scan

import (
	"encoding/binary"
	"fmt"
	"net"
)

// LocalSubnet returns this machine's IPv4 address and subnet mask (for example
// 192.168.1.23/24) from the first network interface that is up.
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
		// An interface usually has several addresses (IPv4 plus one or more
		// IPv6).
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

// Hosts lists every usable host address in ipNet, leaving out the network
// address, the broadcast address and ipNet's own IP. Subnets larger than a /22
// (1,022 hosts) are refused so a big network doesn't turn into a huge sweep.
func Hosts(ipNet *net.IPNet) ([]net.IP, error) {
	ip4 := ipNet.IP.To4()
	// For a /24, ones is 24 and bits is 32, leaving 8 bits for host addresses.
	ones, bits := ipNet.Mask.Size()
	if bits-ones > 10 {
		return nil, fmt.Errorf("subnet /%d too large to sweep", ones)
	}

	// Treating the address as a single 32-bit number turns walking the range
	// into counting: 192.168.1.0 + 1 is 192.168.1.1, and so on.
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
