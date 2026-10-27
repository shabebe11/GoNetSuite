// GoNetSuite is a command-line network toolkit built with Go's standard
// library. Running it starts an interactive shell; type "help" for commands.
package main

import (
    "os"
    "github.com/shabebe11/GoNetSuite/internal/cli"
)

func main() {
    cli.RunShell(os.Stdin, os.Stdout)
}