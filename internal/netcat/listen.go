package netcat

import (
	"context"
	"io"
	"log"
	"net"
	"fmt"
)

func Listen(ctx context.Context, port int, out io.Writer) error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
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