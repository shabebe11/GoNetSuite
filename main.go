package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("\n Enter a command or type 'exit' to quit: ")
        if(!scanner.Scan()) {
            break
        }

        input:= strings.TrimSpace(scanner.Text())
        if input == "" {
            continue
        }

        switch input {
        case "exit":
            fmt.Println("Exiting the program.")
            return
        
        default:
            fmt.Println("Meow")
        }
    }
}