package command 

import (
          "context"
          "fmt"
          "io"
  )

func ListenCommand(ctx context.Context, args []string, in io.Reader, out io.Writer) error {
	fmt.Fprintln(out, "Listening for connections...")
	return nil
}