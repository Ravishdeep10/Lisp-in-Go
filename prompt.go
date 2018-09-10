package main

import (
    "bufio"
    "fmt"
    "os"
    "github.com/alecthomas/participle"
    //"github.com/alecthomas/repr"
)


type CLI struct{
}


func (cli *CLI) prompt() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("go-lisp Version 0.0.0.0.1")
    fmt.Println("Press Ctrl+c to Exit")

    parser, err := participle.Build(&Go_lisp{})

    for{
        fmt.Print("go-lisp> ")
        text, _ := reader.ReadString('\n')

        ini := &Go_lisp{}
        err = parser.ParseString(text, ini)

        if err != nil {
		panic(err)
	    }

        eval := &Evaluator{nil, nil, 0}
        eval.evaluate(ini)
        fmt.Println(eval.Result)

    }
}
