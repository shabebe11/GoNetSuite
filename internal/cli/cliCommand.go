package cli

import (
	"context"
	"io"
	"github.com/shabebe11/GoNetSuite/internal/command"
)

type Command struct {
	Name string
	Usage string
	Run func(ctx context.Context, args []string, in io.Reader, out io.Writer) error
}

var commands = []Command {
	{Name: "scan", Usage: "Scan a range of TCP ports on a host", Run: command.ScanCommand},
	{Name: "sniff", Usage: "Capture and decode packets", Run: command.SniffCommand},
	// {Name: "connect", Usage: "Open a TCP/UDP connection", Run: connectCommand},
	{Name: "listen", Usage: "Listen for TCP/UDP connections", Run: command.ListenCommand},
}