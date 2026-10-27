package command

import (
          "context"
          "fmt"
          "io"
  )

func ScanCommand(ctx context.Context, args []string, in io.Reader, out io.Writer) error {
	fmt.Fprintln(out, "Scanning ports...")
	return nil
}