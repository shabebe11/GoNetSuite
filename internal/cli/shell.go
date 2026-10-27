package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func RunShell(in io.Reader, out io.Writer) {
scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter a command or type 'exit' to quit: ")
		if(!scanner.Scan()) {
			break
		}

		input:= strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		switch input {
			case "scan":
				fmt.Println("Scanning ports...")
			case "exit":
				fmt.Println("\nExiting the program.")
				return
			
			default:
				fmt.Println("\nMeow\n")
		}
    }
}