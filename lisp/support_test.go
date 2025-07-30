package lisp

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/rtbaker/goyali/lexer"
	"github.com/rtbaker/goyali/parser"
)

type SimpleTest struct {
	Code          string
	Expected      string
	ExpectedError string
}

// Support func's for testing
func runExpression(expr string) (string, error) {
	reader := bufio.NewReader(strings.NewReader(expr))
	lex := lexer.NewLexer(reader)
	myParser := parser.NewParser(lex)

	program, err := myParser.ParseProgram()

	if err != nil {
		return "", err
	}

	var b strings.Builder

	for _, expression := range program.Expressions {
		retNode, err := EvaluateNode(expression, program.env, false)

		if err != nil {
			return "", err
		}

		if !IsNil(retNode) {
			b.WriteString(fmt.Sprintf("%s\n", retNode))
		}
	}

	return b.String(), nil
}
