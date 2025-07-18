# goyali

Go Yet Another Lisp Interpreter

Based on https://paulgraham.com/rootsoflisp.html.


### Run Tests ###

`go test -v -coverprofile cover.out  ./...`

Show coverage report:

`go tool cover -html=cover.out`

## Grammar ##

program -> expressions
expressions -> expressions expression | [EMPTY]
expression -> list | atom
list -> '(' expressions ')'
atom -> [SEQUENCE OF LETTERS]



