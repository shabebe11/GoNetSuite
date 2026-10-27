package cli

type command struct {
	Name string
	Usage string
	// Run func(ctx context.Context, args []string, in io.Reader, out io.Writer) error
}

var commands = []command {
	{Name: "scan", Usage: "Scan a range of TCP ports on a host"},
	{Name: "sniff", Usage: "Capture and decode packets"},
	{Name: "connect", Usage: "Open a TCP/UDP connection"},
	{Name: "listen", Usage: "Listen for TCP/UDP connections"},
}