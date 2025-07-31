package lisp

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/rtbaker/goyali/lexer"
)

type SimpleTest struct {
	Code          string
	LoadFromFile  string
	Expected      string
	ExpectedError string
}

// Support func's for testing
func runExpression(expr string) (string, error) {
	reader := bufio.NewReader(strings.NewReader(expr))
	lex := lexer.NewLexer(reader)
	myParser := NewParser(lex)

	program, err := myParser.ParseProgram()

	if err != nil {
		return "", err
	}

	var b strings.Builder

	for index, expression := range program.Expressions {
		retNode, err := EvaluateNode(expression, program.env, false)

		if err != nil {
			return "", err
		}

		if index > 0 {
			b.WriteString("\n")
		}
		b.WriteString(fmt.Sprintf("%s", retNode))
	}

	return b.String(), nil
}
