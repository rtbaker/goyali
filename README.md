# goyali

Go Yet Another Lisp Interpreter

Based on https://paulgraham.com/rootsoflisp.html.

No type system (yet).

## Build ##

`go build -o goyali main.go`

## Use ##

`goyali --lib=./lib`

### Run Tests ###

`go test -v -coverprofile cover.out  ./...`

Show coverage report:

`go tool cover -html=cover.out`

### Todo ###

- Better error handling
- Type system
- Better interactive CLI



