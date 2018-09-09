package main

import (
    "bufio"
    "fmt"
    "os"
)


type CLI struct{
}


func (cli *CLI) prompt() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("go-lisp Version 0.0.0.0.1")
    fmt.Println("Press Ctrl+c to Exit")

    for{
        fmt.Print("go-lisp> ")
        text, _ := reader.ReadString('\n')
        fmt.Printf("No you are a %s\n", text)
    }
}
