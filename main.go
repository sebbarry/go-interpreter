package main 


import (
    "fmt"
    "os"
    "interpreter/repl"
)

func main() {
    fmt.Printf("Interpreter: Start typing here.")
    repl.Start(os.Stdin, os.Stdout)
}
