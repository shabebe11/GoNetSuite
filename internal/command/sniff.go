package command 

import (
          "context"
          "fmt"
          "io"
  )

// SniffCommand will capture and decode packets on a network interface. It is
// a placeholder for now.
func SniffCommand(ctx context.Context, args []string, in io.Reader, out io.Writer) error {
	fmt.Fprintln(out, "Capturing and decoding packets...")
	return nil
}