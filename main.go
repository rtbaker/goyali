package main

import (
	"fmt"
	"os"

	"github.com/rtbaker/goyali/lexer"
	"github.com/rtbaker/goyali/lisp"
)

type LispTest struct {
	Name     string
	Code     string
	Filename string
}

func main() {
	lex := lexer.NewLexer(os.Stdin)

	myParser := lisp.NewParser(lex)

	// Setup top level env/symbol table
	env := lisp.NewEnv(nil)
	env.InitialiseBuiltin()

	fmt.Printf("> ")

	var node lisp.Node
	var err error

	node, err = myParser.GetExpression()

	for {
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}

		// EOF
		if node == nil {
			break
		}

		var resultNode lisp.Node
		resultNode, err = lisp.EvaluateNode(node, env, false)

		if err != nil {
			fmt.Printf("error: %s\n", err)
		}

		fmt.Printf("%s\n> ", resultNode)
		node, err = myParser.GetExpression()
	}

	fmt.Println()
}
