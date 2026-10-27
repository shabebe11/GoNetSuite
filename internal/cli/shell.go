package cli

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"github.com/shabebe11/GoNetSuite/internal/command"
	"context"
)

func RunShell(in io.Reader, out io.Writer) {
	var scanner = bufio.NewScanner(in)
	for {
		fmt.Fprintln(out, "Enter a command or type 'exit' to quit: ")
		if(!scanner.Scan()) {
			break
		}

		input:= strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		inputArgs := strings.Fields(input)
		name, args := inputArgs[0], inputArgs[1:]

		switch name {
			case "scan":
				command.ScanCommand(context.Background(), args, in, out)
			case "sniff":
				command.SniffCommand(context.Background(), args, in, out)
			case "listen":
				command.ListenCommand(context.Background(), args, in, out)
			case "exit":
				fmt.Fprintln(out, "\nExiting the program.")
				return
			
			default:
				fmt.Fprintln(out, "\nUnknown command\n")
		}
    }
}