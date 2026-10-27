package command

import (
          "context"
          "fmt"
          "io"
  )

// ScanCommand will scan a host for open TCP ports. It is a placeholder for now.
func ScanCommand(ctx context.Context, args []string, in io.Reader, out io.Writer) error {
	fmt.Fprintln(out, "Scanning ports...")
	return nil
}