package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/rtbaker/goyali/lexer"
	"github.com/rtbaker/goyali/lisp"
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
		{Name: "Quote op atom", Code: "(quote a)"},
		{Name: "Quote op list", Code: "(quote (a b c))"},
		{Name: "Short Quote op atom", Code: "'a"},
		{Name: "Short Quote op list", Code: "'(a b c)"},
		{Name: "Short Quote op nested list", Code: "'(a b (c))"},
		{Name: "Atom op atom", Code: "(atom a)"},
		{Name: "Atom op quoted atom", Code: "(atom 'a)"},
		{Name: "Atom op list", Code: "(atom (a b c))"},
		{Name: "Atom op quoted list", Code: "(atom '(a b c))"},
		{Name: "atom atom quote", Code: "(atom (atom 'a))"},
		{Name: "atom quote atom atom", Code: "(atom '(atom 'a))"},
		{Name: "eq one arg", Code: "(eq 'a)"},
		{Name: "equals atom atom", Code: "(eq a b)"},
		{Name: "equals qatom qatom same", Code: "(eq 'a 'a)"},
		{Name: "equals qatom qatom not the same", Code: "(eq 'a 'b)"},
		{Name: "Car OP", Code: "(car (a b c))"},
		{Name: "Car OP", Code: "(car '(a b c))"},
		{Name: "Car OP", Code: "(car 'a)"},
		{Name: "Car OP", Code: "(car (car '(a b c)))"},
		{Name: "Cdr OP", Code: "(cdr (a b c))"},
		{Name: "Cdr OP", Code: "(cdr '(a b c))"},
		{Name: "Cdr OP", Code: "(cdr (cdr '(a b c)))"},
		{Name: "Cons OP (bad)", Code: "(cons a (b c d))"},
		{Name: "Cons OP (bad)", Code: "(cons '(a) '(b c d))"},
		{Name: "Cons OP (bad)", Code: "(cons 'a 'b)"},
		{Name: "Cons OP", Code: "(cons 'a '(b c))"},
		{Name: "Cons OP", Code: "(cons 'a (cons 'b (cons 'c '())))"},
		{Name: "Car Cons OP", Code: "(car (cons 'a '(b c)))"},
		{Name: "Cdr Cons OP", Code: "(cdr (cons 'a '(b c)))"},
		{Name: "Cond OP", Code: "(cond ((eq 'a 'b) 'first) ((atom 'a) 'second))"},
		{Name: "Cond OP", Code: "(cond ((eq 'a 'a) 'first) ((atom 'a) 'second))"},
		{Name: "Cond OP", Code: "(cond ((eq 'a 'b) 'first) ((eq 'a 'b) 'second))"},
		{Name: "Lambda", Code: "(lambda (x) (cons x '(b)))"},
		{Name: "Lambda run", Code: "((lambda (x) (cons x '(b))) 'a)"},
		{Name: "Lambda run", Code: "((lambda (x y) (cons x (cdr y))) 'z '(a b c))"},
		{Name: "Lambda run no args", Code: "((lambda () 'a))"},
		/*
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
		*/
	}

	for _, test := range tests {
		//reader := bufio.NewReader(os.Stdin)
		reader := bufio.NewReader(strings.NewReader(test.Code))
		lex := lexer.NewLexer(reader)

		myParser := parser.NewParser(lex)
		program, err := myParser.Parse()

		fmt.Printf("Test: %s\n", test.Name)

		if err != nil {
			fmt.Printf("parse error: %s\n", err)
			return
		}

		/*
			for _, expression := range program.Children() {
				indent := 1

				if listExpr, ok := expression.(lisp.ListNode); ok {
					lisp.WalkTree(
						listExpr,
						func(n lisp.Node) error {
							printSpaces(indent)
							fmt.Printf("%s\n", n)
							return nil
						},
						func() error { indent++; return nil },
						func() error { indent--; return nil },
					)
				} else {
					printSpaces(indent)
					fmt.Printf("%s\n", expression)
				}
			}
		*/

		err = lisp.RunProgram(program)

		if err != nil {
			fmt.Printf("%s\n", err)
		}

		fmt.Println()
	}
}

func printSpaces(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Printf(" ")
	}

	fmt.Printf("- ")
}
