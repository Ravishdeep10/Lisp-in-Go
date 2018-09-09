package main

import (
    "fmt"
    "github.com/alecthomas/participle"
    "github.com/alecthomas/repr"
)

type Go_lisp struct {
    Op *Operator `@@`
    Expr []*Expression `{ @@ }`
}

type Expression struct {
    Digit *int `@Int`
    Op *Operator `| "(" @@`
    Expr []*Expression `{ @@ } ")"`
}

type Operator struct {
    Op *string `"+" | "-" | "*" | "/"`
}

func main() {
    parser, err := participle.Build(&Go_lisp{})
    ini := &Go_lisp{}
    err = parser.ParseString(`
        + 5 (* 2 2)
    `, ini)

    if err != nil {
        fmt.Println(err)
    } else {
        repr.Println(ini, repr.Indent("  "), repr.OmitEmpty(true))
    }
}
