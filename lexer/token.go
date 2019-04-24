package lexer

import (
	"fmt"
)

// TokenCode The type of the token
type TokenCode int

const (
	TokEOF         TokenCode = 1
	TokINT         TokenCode = 2
	TokFLOAT       TokenCode = 3
	TokPLUS        TokenCode = 4
	TokMINUS       TokenCode = 5
	TokMULTIPLY    TokenCode = 6
	TokDIVIDE      TokenCode = 7
	TokOPENBRACE   TokenCode = 8
	TokCLOSEBRACE  TokenCode = 9
	TokSTRING      TokenCode = 10
	TokOPENPARENS  TokenCode = 11
	TokCLOSEPARENS TokenCode = 12
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
	case TokEOF:
		fmt.Println("End of file")
	case TokINT:
		fmt.Printf("Integer: %d\n", tok.IntValue)
	case TokFLOAT:
		fmt.Printf("Float: %f\n", tok.FloatValue)
	case TokPLUS:
		fmt.Println("Plus")
	case TokMINUS:
		fmt.Println("Minus")
	case TokMULTIPLY:
		fmt.Println("Multiply")
	case TokDIVIDE:
		fmt.Println("Divide")
	case TokOPENBRACE:
		fmt.Println("Open brace")
	case TokCLOSEBRACE:
		fmt.Println("Close brace")
	case TokOPENPARENS:
		fmt.Println("Open parens")
	case TokCLOSEPARENS:
		fmt.Println("Close parens")
	case TokSTRING:
		fmt.Printf("String: %s\n", tok.StringValue)
	}
}
