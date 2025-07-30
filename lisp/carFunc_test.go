package lisp

import (
	"fmt"
	"strings"
	"testing"
)

func TestCarOp(t *testing.T) {
	tests := []SimpleTest{
		{Code: "(car '(a b c))", Expected: "a"},
		{Code: "(car '((a b c) d))", Expected: "(a b c)"},
		{Code: "(car 'a)", ExpectedError: "car operator requires a list as its argument"},
	}

	for _, test := range tests {
		output, err := runExpression(test.Code)

		if err != nil {
			if test.ExpectedError == "" {
				t.Errorf("got unexpected error: %s", err)
			}

			errStr := fmt.Sprintf("%s", err)

			if !strings.HasPrefix(errStr, test.ExpectedError) {
				t.Errorf("error not correct, expected \"%s\" got \"%s\"", test.ExpectedError, errStr)
			}

			continue
		}

		t.Logf("*** %s -> %s", test.Code, output)

		if output != test.Expected {
			t.Errorf("expression output incorrect, expected \"%s\" got \"%s\"", test.Expected, output)
		}
	}
}
