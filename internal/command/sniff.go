package command 

import (
          "context"
          "fmt"
          "io"
  )

func SniffCommand(ctx context.Context, args []string, in io.Reader, out io.Writer) error {
	fmt.Fprintln(out, "Capturing and decoding packets...")
	return nil
}