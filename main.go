package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/rtbaker/goyali/lexer"
	"github.com/rtbaker/goyali/lisp"
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
		{Name: "2 Atom Lists", Code: "(foo bar hello) (foo bar hello)"},
		{Name: "Nested list", Code: "(a b (c d) e)"},
		{Name: "Quote op atom", Code: "(quote a)"},
		{Name: "Quote op list", Code: "(quote (a b c))"},
		{Name: "Short Quote op atom", Code: "'a"},
		{Name: "Short Quote op list", Code: "'(a b c)"},
		{Name: "Short Quote op nested list", Code: "'(a b (c))"},
		{Name: "Atom op atom", Code: "(atom a)"},
		{Name: "Atom op list", Code: "(atom (a b c))"},
		{Name: "Equals op atom", Code: "(eq a b)"},
		{Name: "Equals op empty lists", Code: "(eq () ())"},
		{Name: "Equals op empty list and atom", Code: "(eq a (b c))"},
		{Name: "Car OP", Code: "(car (a b c))"},
		{Name: "Cdr OP", Code: "(cdr (a b c))"},
		{Name: "Cons OP", Code: "(cons a (b c d))"},
		{Name: "Cond OP", Code: "(cond ((eq a b) first) ((atom a) second))"},
		{Name: "Cond OP 2", Code: "(cond (y 't) ('t '()))"},
		{Name: "Lambda", Code: "(lambda (x) (cons x (b)))"},
		{Name: "Lambda with args", Code: "((lambda (x) (cons x (b))) a)"},
		{Name: "Defun op", Code: "(defun subst (a b c) (cons x (b)))"},
		{Name: "Label", Code: "(label f (lambda (x y z) (cons x(b))))"},
		{Name: "Bad quote op (2 args)", Code: "(quote a b)"},
		{Name: "Bad atom op (2 args)", Code: "(atom a b)"},
		{Name: "Bad equals op (1 args)", Code: "(eq a)"},
		{Name: "Bad equals op (3 args)", Code: "(eq a b c)"},
		{Name: "Bad car op (2 args)", Code: "(car a b)"},
		{Name: "Bad cdr op (2 args)", Code: "(cdr a b)"},
		{Name: "Bad cons op (1 args)", Code: "(cons a)"},
		{Name: "Bad Cond OP", Code: "(cond ((eq a b) ) ((atom a) second))"},
		{Name: "Bad Label (1 arg)", Code: "(label (lambda (x y z) (cons x(b))))"},
		{Name: "Bad Label (non atom first arg)", Code: "(label (f) (lambda (x y z) (cons x(b))))"},
		{Name: "Bad lambda (1 arg)", Code: "(lambda (cons x (b)))"},
		{Name: "Bad lambda (non atom arg)", Code: "(lambda (a (a)) (cons x (b)))"},
		{Name: "Bad defun (name not atom)", Code: "(defun (a b) (a b) 'a)"},
		{Name: "Bad defun (non atom arg)", Code: "(defun name (a (b)) 'a)"},
	}

	for _, test := range tests {
		//reader := bufio.NewReader(os.Stdin)
		reader := bufio.NewReader(strings.NewReader(test.Code))
		lex := lexer.NewLexer(reader)

		myParser := lisp.NewParser(lex)
		node, err := myParser.Parse()

		if err != nil {
			fmt.Printf("%s case error: %s", test.Name, err)
			return
		}

		fmt.Printf("Test: %s\n", test.Name)

		err = lisp.SyntaxCheckTree(node)
		if err != nil {
			fmt.Printf("Syntax Check error: %s\n\n", err)
			continue
		}

		indent := 1
		lisp.WalkTree(
			node,
			func(n lisp.Node) error {
				printSpaces(indent)
				fmt.Printf("%s\n", n)
				return nil
			},
			func() error { indent++; return nil },
			func() error { indent--; return nil },
		)

		fmt.Println()
	}
}

func printSpaces(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Printf(" ")
	}

	fmt.Printf("- ")
}
