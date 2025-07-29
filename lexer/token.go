package lexer

import (
	"fmt"
	"strings"
)

// TokenCode The type of the token
type TokenCode int

const (
	EOF TokenCode = iota
	ATOM
	OPENPARENS
	CLOSEPARENS
	SHORTQUOTE
)

// Token A single token from the lexer
type Token struct {
	Code TokenCode
	// IntValue    int64
	// FloatValue  float64
	// StringValue string
	Value    string
	Line     int
	Position int
}

func (tok *Token) String() string {
	var name string
	var val string

	switch tok.Code {
	case EOF:
		name = "EOF"
	case SHORTQUOTE:
		name = "SHORTQUOTE"
	case ATOM:
		name = "ATOM"
		val = tok.Value
	case OPENPARENS:
		name = "OPENPARENS"
	case CLOSEPARENS:
		name = "CLOSEPARENS"
	}

	if val != "" {
		return fmt.Sprintf("%s: %s\t\t\t(Line %d, position %d)", name, val, tok.Line, tok.Position)
	}

	return fmt.Sprintf("%s\t\t\t(Line %d, position %d)", name, tok.Line, tok.Position)
}

func StringToTokenType(val string) TokenCode {
	val = strings.ToUpper(val)

	switch val {
	case "SHORTQUOTE":
		return SHORTQUOTE
	}

	return ATOM
}
