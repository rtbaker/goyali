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
		{"+", &Token{Code: PLUS}},
		{"-", &Token{Code: MINUS}},
		{"/", &Token{Code: DIVIDE}},
		{"*", &Token{Code: MULTIPLY}},
		{"{", &Token{Code: OPENBRACE}},
		{"}", &Token{Code: CLOSEBRACE}},
		{"(", &Token{Code: OPENPARENS}},
		{")", &Token{Code: CLOSEPARENS}},
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

func TestEmptyString(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("      "))
	lex := NewLexer(reader)

	token, err := lex.GetToken()

	if err != nil {
		t.Errorf("Lexer GetToken returned an error: %s", err)
	}

	expected := Token{Code: EOF}

	if *token != expected {
		t.Errorf("Lexer GetToken dod not return EOF Token: got %d ", token.Code)
	}
}

func TestInt(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("123"))
	lex := NewLexer(reader)

	token, err := lex.GetToken()

	if err != nil {
		t.Errorf("Lexer GetToken returned an error: %s", err)
	}

	expected := Token{Code: INT, IntValue: 123}

	if *token != expected {
		if token.Code != expected.Code {
			t.Errorf("Lexer GetToken did not return INT Token: got %d ", token.Code)
		}

		if token.IntValue != expected.IntValue {
			t.Errorf("Lexer did not return correct in value (expected 123) got %d", token.IntValue)
		}
	}
}

func TestFloat(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("123.123"))
	lex := NewLexer(reader)

	token, err := lex.GetToken()

	if err != nil {
		t.Errorf("Lexer GetToken returned an error: %s", err)
	}

	expected := Token{Code: FLOAT, FloatValue: 123.123}

	if *token != expected {
		if token.Code != expected.Code {
			t.Errorf("Lexer GetToken did not return Float Token: got %d ", token.Code)
		}

		if token.IntValue != expected.IntValue {
			t.Errorf("Lexer did not return correct in value (expected 123.123) got %f", token.FloatValue)
		}
	}
}

func TestBadFloat(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("123.123.123"))
	lex := NewLexer(reader)

	_, err := lex.GetToken()

	if err == nil {
		t.Errorf("Lexer GetToken should have returned an error but did not")
	} else if err.Error() != "Double '.' in number" {
		t.Errorf("Lexer GetToken did not return the correct error string: %s", err)
	}
}

func TestString(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("somestring"))
	lex := NewLexer(reader)

	token, err := lex.GetToken()

	if err != nil {
		t.Errorf("Lexer GetToken returned an error: %s", err)
	}

	expected := Token{Code: STRING, StringValue: "somestring"}

	if *token != expected {
		if token.Code != expected.Code {
			t.Errorf("Lexer GetToken did not return String Token: got %d ", token.Code)
		}

		if token.StringValue != expected.StringValue {
			t.Errorf("Lexer did not return correct in value (expected 'somestring') got %s", token.StringValue)
		}
	}
}

func TestEOFAfterMultipleTokens(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("somestring and another 123 + 45"))
	lex := NewLexer(reader)

	var token *Token
	var err error

	for token, err = lex.GetToken(); err == nil && token.Code != EOF; {
		// Get next
		token, err = lex.GetToken()
	}

	if err != nil {
		t.Errorf("Reading multiple tokens returned an error: %s", err)
	}

	// Last token should be EOF
	expected := Token{Code: EOF}

	if *token != expected {
		t.Errorf("Lexer GetToken dod not return EOF Token: got %d ", token.Code)
	}
}

func TestChar(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("@"))
	lex := NewLexer(reader)

	_, err := lex.GetToken()

	if err == nil {
		t.Errorf("Lexer GetToken should have returned an error but did not")
	} else if err.Error() != "Unrecognised token: '@'" {
		t.Errorf("Lexer GetToken did not return the correct error string: %s", err)
	}
}
