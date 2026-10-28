package command

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/shabebe11/GoNetSuite/internal/netcat"
)

// ListenCommand handles "listen <port>". It checks that the port is a number
// from 1 to 65535, then listens on it until stopped.
func ListenCommand(ctx context.Context, args []string, in io.Reader, out io.Writer) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: listen <port>\n")
	}
	port, err := strconv.Atoi(args[0])
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("> invalid port number: %s", args[0])
	}
	workerCtx, workerCancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer workerCancel()

	go netcat.Listen(workerCtx, port, out)
	<-workerCtx.Done()
	return nil
}
