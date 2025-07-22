package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/rtbaker/goyali/lexer"
	"github.com/rtbaker/goyali/parser"
)

type LispTest struct {
	Name string
	Code string
}

func main() {
	tests := []LispTest{
		{Name: "Simple Atom", Code: "foo"},
		{Name: "Empty List", Code: "()"},
		{Name: "Atom List", Code: "(foo bar hello)"},
	}

	for _, test := range tests {
		//reader := bufio.NewReader(os.Stdin)
		reader := bufio.NewReader(strings.NewReader(test.Code))
		lex := lexer.NewLexer(reader)

		parser := parser.NewParser(lex)
		node, err := parser.Parse()

		if err != nil {
			fmt.Printf("%s case error: %s", test.Name, err)
			return
		}

		fmt.Printf("Test: %s\n", test.Name)
		printTree(node, 1)
		fmt.Println()
	}
}

func printTree(n parser.Node, indent int) {
	printSpaces(indent)
	fmt.Printf("%s\n", n)

	for _, n := range n.Children() {
		printTree(n, indent+1)
	}
}

func printSpaces(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Printf(" ")
	}

	fmt.Printf("- ")
}
