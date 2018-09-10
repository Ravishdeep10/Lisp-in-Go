
# Lisp in Go

Implementing the Lisp programming language in Go.
To run :
```bash
go build .
./lisp_in_go
```

For now it can only evaluate polish notation equations but will add s-expressions, variables, and functions in the foreseeable future.

```bash
go-lisp Version 0.0.0.0.1
Press Ctrl+c to Exit
go-lisp> + 5 6
11
go-lisp> - (* 10 10) (+ 1 1 1)
97
```
Used "alecthomas/participle" library as a Go implemented parser.
