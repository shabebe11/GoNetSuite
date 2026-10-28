// Package cli implements the interactive GoNetSuite shell.
package cli

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/shabebe11/GoNetSuite/internal/command"
)

// RunShell reads commands line by line from in, runs them, and writes all
// output and errors to out. It returns on "exit" or when in reaches EOF.
func RunShell(in io.Reader, out io.Writer) {
	var scanner = bufio.NewScanner(in)
	for {
		fmt.Fprintln(out, "> Enter a command or type 'exit' to quit: ")
		if !scanner.Scan() {
			break
		}
		if err := scanner.Err(); err != nil {
			// Catch and handle the error here
			log.Fatal(err)
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		inputArgs := strings.Fields(input)
		name, args := inputArgs[0], inputArgs[1:]

		//background context for the worker goroutine, which will be canceled on SIGINT or SIGTERM (Ctrl+C)
		workerCtx, workerCancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

		var err error
		switch name {
		case "scan":
			err = command.ScanCommand(workerCtx, args, in, out)
		case "sniff":
			err = command.SniffCommand(workerCtx, args, in, out)
		case "listen":
			err = command.ListenCommand(workerCtx, args, in, out)
		case "help":
			err = command.HelpCommand(workerCtx, args, in, out)
		case "exit":
			fmt.Fprintln(out, "\n> Exiting the program.")
			return

		default:
			fmt.Fprintln(out, "\n> Unknown command\n")
		}
		defer workerCancel()

		if err != nil {
			fmt.Fprintf(out, "> Error executing command: %v\n", err)
		}
	}
}
