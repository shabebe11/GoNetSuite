package command

import (
	"context"
	"fmt"
	"io"
)

func HelpCommand(ctx context.Context, args []string, in io.Reader, out io.Writer) error {
	fmt.Fprintln(out, "-------------------------------------")
	fmt.Fprintln(out, "Available commands:")
	fmt.Fprintln(out, "  scan   - Scan ports")
	fmt.Fprintln(out, "  sniff  - Capture and decode packets")
	fmt.Fprintln(out, "  listen - Listen for connections")
	fmt.Fprintln(out, "  help   - Show this help message")
	fmt.Fprintln(out, "  exit   - Exit the program")
	fmt.Fprintln(out, "-------------------------------------")

	return nil;
}