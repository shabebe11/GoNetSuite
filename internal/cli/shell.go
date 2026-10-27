// Package cli implements the interactive GoNetSuite shell.
package cli

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"github.com/shabebe11/GoNetSuite/internal/command"
	"context"
)

// RunShell reads commands line by line from in, runs them, and writes all
// output and errors to out. It returns on "exit" or when in reaches EOF.
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

		var err error
		switch name {
			case "scan":
				err = command.ScanCommand(context.Background(), args, in, out)
			case "sniff":
				err = command.SniffCommand(context.Background(), args, in, out)
			case "listen":
				err = command.ListenCommand(context.Background(), args, in, out)
			case "help":
				err = command.HelpCommand(context.Background(), args, in, out)
			case "exit":
				fmt.Fprintln(out, "\nExiting the program.")
				return
			
			default:
				fmt.Fprintln(out, "\nUnknown command\n")
		}
		if err != nil {
			fmt.Fprintf(out, "Error executing command: %v\n", err)
		}
    }
}