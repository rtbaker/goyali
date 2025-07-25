package lexer

import (
	"fmt"
	"strings"
)

// TokenCode The type of the token
type TokenCode int

const (
	EOF TokenCode = iota
	//	INT
	//	FLOAT
	//	PLUS
	//	MINUS
	//	MULTIPLY
	//	DIVIDE
	//	OPENBRACE
	//	CLOSEBRACE
	//	STRING
	ATOM
	OPENPARENS
	CLOSEPARENS
	// COMMA
	QUOTE
	SHORTQUOTE
	ATOMOPERATOR
	EQUALS
	CAR
	CDR
	CONS
	COND
	LAMBDA
	LABEL
	DEFUN
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
	case QUOTE:
		name = "QUOTE"
	case SHORTQUOTE:
		name = "SHORTQUOTE"
	case ATOMOPERATOR:
		name = "ATOMOPERATOR"
	case EQUALS:
		name = "EQUALS"
	case ATOM:
		name = "ATOM"
		val = tok.Value
	case CAR:
		name = "CAR"
	case CDR:
		name = "CDR"
	case CONS:
		name = "CONS"
	case COND:
		name = "COND"
	case LAMBDA:
		name = "LAMBDA"
	case LABEL:
		name = "LABEL"
	case DEFUN:
		name = "DEFUN"
	case OPENPARENS:
		name = "OPENPARENS"
	case CLOSEPARENS:
		name = "CLOSEPARENS"
		//	case INT:
		//		name = "INT"
		//		val = fmt.Sprintf("%d", tok.IntValue)
		//	case FLOAT:
		//		name = "FLOAT"
		//		val = fmt.Sprintf("%f", tok.FloatValue)
		//	case PLUS:
		//		name = "PLUS"
		//	case MINUS:
		//		name = "MINUS"
		//	case MULTIPLY:
		//		name = "MULTIPLY"
		//	case DIVIDE:
		//		name = "DIVIDE"
		//	case OPENBRACE:
		//		name = "OPENBRACE"
		//	case CLOSEBRACE:
		//		name = "CLOSEBRACE"
		//	case COMMA:
		//		name = "COMMA"
		//	case STRING:
		//		name = "STRING"
		//		val = tok.StringValue
	}

	if val != "" {
		return fmt.Sprintf("%s: %s\t\t\t(Line %d, position %d)", name, val, tok.Line, tok.Position)
	}

	return fmt.Sprintf("%s\t\t\t(Line %d, position %d)", name, tok.Line, tok.Position)
}

func StringToTokenType(val string) TokenCode {
	val = strings.ToUpper(val)

	switch val {

	case "QUOTE":
		return QUOTE
	case "SHORTQUOTE":
		return SHORTQUOTE
	case "ATOM":
		return ATOMOPERATOR
	case "EQ":
		return EQUALS
	case "CAR":
		return CAR
	case "CDR":
		return CDR
	case "CONS":
		return CONS
	case "COND":
		return COND
	case "LAMBDA":
		return LAMBDA
	case "LABEL":
		return LABEL
	case "DEFUN":
		return DEFUN
	}

	return ATOM
}
