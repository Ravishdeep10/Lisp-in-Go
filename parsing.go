package main


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
    OpAdd *string `@"+"`
    OpSub *string `| @"-"`
    OpMul *string `| @"*"`
    OpDiv *string `| @"/"`
}
