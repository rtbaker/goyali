package lisp

import (
	"strings"
	"testing"
)

func TestFirstArgNoValueFunction(t *testing.T) {
	_, err := runExpression("(notAFunc 'a 'b)")

	if err == nil {
		t.Errorf("should return an error but doesn't")
		return
	}

	expected := "Atom \"notAFunc\" has no value"
	if !strings.HasPrefix(err.Error(), expected) {
		t.Errorf("got error: \"%s\" but expected \"%s\"", err.Error(), expected)
	}
}
