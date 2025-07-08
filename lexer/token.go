package lexer

import (
	"fmt"
)

// TokenCode The type of the token
type TokenCode int

const (
	EOF TokenCode = iota
	INT
	FLOAT
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	OPENBRACE
	CLOSEBRACE
	STRING
	OPENPARENS
	CLOSEPARENS
	COMMA
)

// Token A single token from the lexer
type Token struct {
	Code        TokenCode
	IntValue    int64
	FloatValue  float64
	StringValue string
	Line        int
	Position    int
}

func (tok *Token) String() string {
	var name string
	var val string

	switch tok.Code {
	case EOF:
		name = "EOF"
	case INT:
		name = "INT"
		val = fmt.Sprintf("%d", tok.IntValue)
	case FLOAT:
		name = "FLOAT"
		val = fmt.Sprintf("%f", tok.FloatValue)
	case PLUS:
		name = "PLUS"
	case MINUS:
		name = "MINUS"
	case MULTIPLY:
		name = "MULTIPLY"
	case DIVIDE:
		name = "DIVIDE"
	case OPENBRACE:
		name = "OPENBRACE"
	case CLOSEBRACE:
		name = "CLOSEBRACE"
	case OPENPARENS:
		name = "OPENPARENS"
	case CLOSEPARENS:
		name = "CLOSEPARENS"
	case COMMA:
		name = "COMMA"
	case STRING:
		name = "STRING"
		val = tok.StringValue
	}

	if val != "" {
		return fmt.Sprintf("%s: %s\t\t\t(Line %d, position %d)", name, val, tok.Line, tok.Position)
	}

	return fmt.Sprintf("%s\t\t\t(Line %d, position %d)", name, tok.Line, tok.Position)
}

// PrintToken Pretty print the token for debugging
func (tok *Token) PrintToken() {
	switch tok.Code {
	case EOF:
		fmt.Println("End of file")
	case INT:
		fmt.Printf("Integer: %d\n", tok.IntValue)
	case FLOAT:
		fmt.Printf("Float: %f\n", tok.FloatValue)
	case PLUS:
		fmt.Println("Plus")
	case MINUS:
		fmt.Println("Minus")
	case MULTIPLY:
		fmt.Println("Multiply")
	case DIVIDE:
		fmt.Println("Divide")
	case OPENBRACE:
		fmt.Println("Open brace")
	case CLOSEBRACE:
		fmt.Println("Close brace")
	case OPENPARENS:
		fmt.Println("Open parens")
	case CLOSEPARENS:
		fmt.Println("Close parens")
	case STRING:
		fmt.Printf("String: %s\n", tok.StringValue)
	}
}
