package lexer

import (
	"bufio"
	"strings"
	"testing"
)

func TestSingleCharSymbols(t *testing.T) {
	tables := []struct {
		symbol string
		tok    *Token
	}{
		{"+", &Token{Code: TokPLUS}},
		{"-", &Token{Code: TokMINUS}},
		{"/", &Token{Code: TokDIVIDE}},
		{"*", &Token{Code: TokMULTIPLY}},
		{"{", &Token{Code: TokOPENBRACE}},
		{"}", &Token{Code: TokCLOSEBRACE}},
		{"(", &Token{Code: TokOPENPARENS}},
		{")", &Token{Code: TokCLOSEPARENS}},
	}

	for _, table := range tables {
		reader := bufio.NewReader(strings.NewReader(table.symbol))
		lex := NewLexer(reader)

		token, err := lex.GetToken()

		if err != nil {
			t.Errorf("Lexer GetToken returned an error: %s", err)
		}

		if *token != *table.tok {
			t.Errorf("Lexer GetToken returned the wrong token: expected %d got %d ", table.tok.Code, token.Code)
		}
	}
}
