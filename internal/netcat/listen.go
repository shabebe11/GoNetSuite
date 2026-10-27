// Package netcat implements simple TCP tools in the style of the nc utility.
package netcat

import (
	"context"
	"io"
	"log"
	"net"
	"fmt"
)

// Listen accepts TCP connections on port, on every network interface, and
// copies everything each client sends to out. Each client is handled in its
// own goroutine. Listen blocks until ctx is cancelled or accepting fails.
func Listen(ctx context.Context, port int, out io.Writer) error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	// Accept ignores ctx, so closing the listener is what unblocks it.
	go func() {
          <-ctx.Done()
          l.Close()
  }()
	for{
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go func (c net.Conn) {
			io.Copy(out, c)
			c.Close()
		}(conn)
	}
}