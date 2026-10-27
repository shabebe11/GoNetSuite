package main

import (
    "os"
    "github.com/shabebe11/GoNetSuite/internal/cli"
)

func main() {
    cli.RunShell(os.Stdin, os.Stdout)
}