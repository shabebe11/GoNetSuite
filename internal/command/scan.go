package command

import (
	"context"
	"io"

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

	return nil
}
