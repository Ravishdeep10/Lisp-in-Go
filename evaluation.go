package main

import "fmt"

type Evaluator struct {
    Expressions []*int
    Operator    *string
    Result     int
}

func (eval *Evaluator) whichOperator(op *Operator) {
    if op.OpAdd != nil {
        eval.Operator = op.OpAdd
    }

    if op.OpSub != nil {
        eval.Operator = op.OpSub
    }

    if op.OpMul != nil {
        eval.Operator = op.OpMul
    }

    if op.OpDiv != nil {
        eval.Operator = op.OpDiv
    }
}

func (eval *Evaluator) evaluate(ini *Go_lisp) {

    // If user doesn't correctly give a correct expression
    if ini.Op == nil || ini.Expr == nil{
        fmt.Println("Error: There is no expression to evaluate")
        return
    }

    if len(ini.Expr) == 1 && ini.Expr[0].Digit != nil {
        fmt.Println("Error: There is no expression to evaluate")
        return
    }

    eval.evalutePolish(ini.Op, ini.Expr)
}

func (eval *Evaluator) evalutePolish(op *Operator, expr []*Expression) {

    eval.whichOperator(op)

    for i := 0; i < len(expr); i++ {
        if expr[i].Digit != nil {
            eval.Expressions = append(eval.Expressions, expr[i].Digit)
        } else {
            newEval := &Evaluator{nil, nil, 0}
            newEval.evalutePolish(expr[i].Op, expr[i].Expr)
            digit := newEval.Result
            eval.Expressions = append(eval.Expressions, &digit)
        }
    }

    eval.collapse()

}


func (eval *Evaluator) collapse() {
    var result int = *eval.Expressions[0]

    for i := 1; i < len(eval.Expressions); i++ {
        if *eval.Operator == "+" {
           result += *eval.Expressions[i]
        }
        if *eval.Operator == "-" {
           result -= *eval.Expressions[i]
        }
        if *eval.Operator == "*" {
           result *= *eval.Expressions[i]
        }
        if *eval.Operator == "/" {
           result /= *eval.Expressions[i]
        }
    }

    eval.Result = result

}
