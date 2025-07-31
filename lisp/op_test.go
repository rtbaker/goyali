package lisp

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// Run single op expressions
func TestQuoteOp(t *testing.T) {
	tests := []SimpleTest{
		{Code: "(quote)", ExpectedError: "quote operator requires only 1 argument"},
		{Code: "(quote a b)", ExpectedError: "quote operator requires only 1 argument"},
		{Code: "(quote a)", Expected: "a"},
		{Code: "(quote (a b c))", Expected: "(a b c)"},
		{Code: "'a", Expected: "a"},
		{Code: "'(a b c)", Expected: "(a b c)"},
		{Code: "'(a (b d) c)", Expected: "(a (b d) c)"},
	}

	runTests(tests, t)
}

func TestEqualsOp(t *testing.T) {
	tests := []SimpleTest{
		{Code: "(eq 'a 'a)", Expected: "t"},
		{Code: "(eq 'a 'b)", Expected: "()"},
		{Code: "(eq 'a)", ExpectedError: "equals operator requires 2 arguments"},
		{Code: "(eq () ())", Expected: "t"},
		{Code: "(eq '() '())", Expected: "t"},
	}

	runTests(tests, t)
}

func TestAtomOp(t *testing.T) {
	tests := []SimpleTest{
		{Code: "(atom a)", ExpectedError: "Atom \"a\" has no value"},
		{Code: "(atom 'a)", Expected: "t"},
		{Code: "(atom (a b c))", ExpectedError: "Atom \"a\" has no value"},
		{Code: "(atom '(a b c))", Expected: "()"},
		{Code: "(atom ())", Expected: "t"},
		{Code: "(atom '())", Expected: "t"},
		{Code: "(atom 'a 'b 'c)", ExpectedError: "atom operator requires only 1 argument"},
	}

	runTests(tests, t)
}

func TestCarOp(t *testing.T) {
	tests := []SimpleTest{
		{Code: "(car '(a b c))", Expected: "a"},
		{Code: "(car '((a b c) d))", Expected: "(a b c)"},
		{Code: "(car '())", Expected: "()"},
		{Code: "(car 'a)", ExpectedError: "car operator requires a list as its argument"},
		{Code: "(car '(a b c) '(a b c))", ExpectedError: "car operator requires only 1 argument"},
	}

	runTests(tests, t)
}

func TestCdrOp(t *testing.T) {
	tests := []SimpleTest{
		{Code: "(cdr '(a b c))", Expected: "(b c)"},
		{Code: "(cdr '((a b c) d))", Expected: "(d)"},
		{Code: "(cdr 'a)", ExpectedError: "cdr operator requires a list as its argument"},
		{Code: "(cdr '(a b c) '(a b c))", ExpectedError: "cdr operator requires only 1 argument"},
	}

	runTests(tests, t)
}

func TestCondOp(t *testing.T) {
	tests := []SimpleTest{
		{Code: "(cond a)", ExpectedError: "argument to cond must be a list"},
		{Code: "(cond '(t))", Expected: "NIL"}, // Not sure if this is correct, clisp claims "quote" not a variable
		{Code: "(cond ((eq 'a 'b) 'first) ('t 'second))", Expected: "second"},
		{Code: "(cond ((eq 'a 'b) 'first) ('t 'second) ('t 'third))", Expected: "second"},
		{Code: "(cond ((eq 'a 'b) 'first) ((eq 'b 'c) 'second) ((eq 'c 'c) 'third))", Expected: "third"},
		{Code: "(cond ((eq 'a 'a) 'first) ('t 'second))", Expected: "first"},
		{Code: "(cond ((eq a 'a) 'first) ('t 'second))", ExpectedError: "Atom \"a\" has no value"},
	}

	runTests(tests, t)
}

func TestConsOp(t *testing.T) {
	tests := []SimpleTest{
		{Code: "(cons 'a)", ExpectedError: "cons operator requires 2 arguments"},
		{Code: "(cons 'a 'b)", ExpectedError: "cons operator requires a list for the second argument"},
		{Code: "(cons 'a ())", Expected: "(a)"},
		{Code: "(cons 'a '())", Expected: "(a)"},
		{Code: "(cons 'a '(b))", Expected: "(a b)"},
		{Code: "(cons 'a '(b c))", Expected: "(a b c)"},
		{Code: "(cons 'a '(b (c d)))", Expected: "(a b (c d))"},
		{Code: "(cons 'a (cdr '(b c d)))", Expected: "(a c d)"},
		{Code: "(cons (car '(b c d)) '(a))", Expected: "(b a)"},
		{Code: "(cons 'a (cons 'b (cons 'c '())))", Expected: "(a b c)"},
	}

	runTests(tests, t)
}

func TestLambdaOp(t *testing.T) {
	tests := []SimpleTest{
		{Code: "(lambda (x) (cons x '(b)))", Expected: "Function"},
		{Code: "(lambda (x))", ExpectedError: "lambda operator requires 2 arguments"},
		{Code: "(lambda ((x)) (cons x '(b)))", ExpectedError: "function arguments must be atoms"},
		// run lambda's
		{Code: "((lambda () 'a))", Expected: "a"},
		{Code: "((lambda (x y) (cons x (cdr y))) 'z '(a b c))", Expected: "(z b c)"},
		{Code: "((lambda (f) (f '(b c))) (lambda (x) (cons 'a x)))", Expected: "(a b c)"}, // one lambda as the param to another
	}

	runTests(tests, t)
}

func TestLabelOp(t *testing.T) {
	tests := []SimpleTest{
		{Code: "(label newFunc)", ExpectedError: "label operator requires 2 arguments"},
		{
			Code:          "(label '(a) (lambda (x y z) (cond ((atom z) ((cond (eq z y) x) ('t z))) ('t (cons (subst x y (car z)) (subst x y (cdr z)))))))",
			ExpectedError: "label op expects first argument to be an atom"},
		{Code: "(label subst (lambda (x y z) (cond ((atom z) ((cond (eq z y) x) ('t z))) ('t (cons (subst x y (car z)) (subst x y (cdr z)))))))", Expected: "NIL"},
		// Run label
		{LoadFromFile: "testCode/label.lisp", Expected: "NIL\n(a m (a m c) d)"},
	}

	runTests(tests, t)
}

func TestDefunOp(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	t.Logf("Current test directory: %s", filepath.Dir(filename))

	tests := []SimpleTest{
		{Code: "(defun '(a) (a) (v))", ExpectedError: "defun op expects first argument to be an atom"},
		{Code: "(defun newFunc)", ExpectedError: "defun operator requires 3 arguments"},
		{Code: "(defun newFunc (a))", ExpectedError: "defun operator requires 3 arguments"},
		{LoadFromFile: "testCode/defun.lisp", Expected: "NIL\n(a m (a m c) d)"},
	}

	runTests(tests, t)
}

func runTests(tests []SimpleTest, t *testing.T) {
	for index, test := range tests {
		code, err := getTestCode(test)

		if err != nil {
			t.Errorf("error getting test code: %s", err)
			continue
		}

		output, err := runExpression(code)

		if err != nil {
			if test.ExpectedError == "" {
				t.Errorf("test %d: got unexpected error: %s", index, err)
			}

			errStr := fmt.Sprintf("%s", err)

			if !strings.HasPrefix(errStr, test.ExpectedError) {
				t.Errorf("test %d: error not correct, expected \"%s\" got \"%s\"", index, test.ExpectedError, errStr)
			}

			continue
		} else if test.ExpectedError != "" {
			t.Errorf("test %d: expected an error but did not get one", index)
			continue
		}

		t.Logf("*** %s -> %s", code, output)

		if output != test.Expected {
			t.Errorf("test %d: expression output incorrect, expected \"%s\" got \"%s\"", index, test.Expected, output)
		}
	}
}

func getTestCode(test SimpleTest) (string, error) {
	if test.Code != "" {
		return test.Code, nil
	}

	if test.LoadFromFile != "" {
		_, filename, _, _ := runtime.Caller(0)
		dir := filepath.Dir(filename)

		filePath := filepath.Join(dir, test.LoadFromFile)

		dat, err := os.ReadFile(filePath)

		if err != nil {
			return "", err
		}

		return string(dat), nil
	}

	return "", fmt.Errorf("no code supplied")
}
