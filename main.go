package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rtbaker/goyali/lexer"
	"github.com/rtbaker/goyali/lisp"
)

type LispTest struct {
	Name     string
	Code     string
	Filename string
}

func main() {
	// Setup top level env/symbol table
	env := lisp.NewEnv(nil)
	env.InitialiseBuiltin()

	fmt.Printf("> ")

	reader := bufio.NewReader(os.Stdin)

	var node lisp.Node
	var err error

	//node, err = myParser.GetExpression()

	var builder strings.Builder

	for {
		text, _ := reader.ReadString('\n')
		builder.WriteString(text)

		lex := lexer.NewLexer(strings.NewReader(builder.String()))

		myParser := lisp.NewParser(lex)
		node, err = myParser.GetExpression()

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

		builder.Reset()
	}

	fmt.Println()
}
