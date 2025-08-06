package lisp

import (
	"fmt"
	"strings"
)

type LispError struct {
	msg      string
	line     int
	position int
	wrapped  error
}

func NewLispError(msg string, line int, position int, wrapped error) *LispError {
	return &LispError{
		msg:      msg,
		line:     line,
		position: position,
		wrapped:  wrapped,
	}
}

func NewSimpleLispError(msg string) *LispError {
	return &LispError{
		msg: msg,
	}
}

func (e *LispError) Error() string {
	var b strings.Builder
	space := ""

	if e.wrapped != nil {
		b.WriteString(e.wrapped.Error())
		space = "\n "
	}

	b.WriteString(fmt.Sprintf("%s%s", space, e.msg))

	if e.line != 0 {
		b.WriteString(fmt.Sprintf(" (line: %d, position %d)", e.line, e.position))
	}

	return b.String()
}

func (e *LispError) Unwrap() error {
	return e.wrapped
}
