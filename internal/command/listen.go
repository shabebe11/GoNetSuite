package command 

import (
	"context"
	"fmt"
	"io"
	"github.com/shabebe11/GoNetSuite/internal/netcat"
	"strconv"
  )

// ListenCommand handles "listen <port>". It checks that the port is a number
// from 1 to 65535, then listens on it until stopped.
func ListenCommand(ctx context.Context, args []string, in io.Reader, out io.Writer) error {
	if len(args) < 1 {
          return fmt.Errorf("usage: listen <port>")
  	}
	port, err := strconv.Atoi(args[0])
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("invalid port number: %s", args[0])
	}
	return netcat.Listen(ctx, port, out)
}