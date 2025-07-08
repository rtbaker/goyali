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
)

// Token A single token from the lexer
type Token struct {
	Code        TokenCode
	IntValue    int64
	FloatValue  float64
	StringValue string
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
