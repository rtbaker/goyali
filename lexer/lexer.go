package lexer

import (
	"bufio"
	"fmt"
	"unicode"
)

// Lexer Simple lex
type Lexer struct {
	in          *bufio.Reader
	currentRune rune
	eof         bool
	pos         int
	line        int
}

// NewLexer Create a new Lexer
func NewLexer(input *bufio.Reader) *Lexer {
	l := new(Lexer)

	l.in = input
	l.currentRune = ' '
	l.eof = false
	l.pos = 1
	l.line = 1

	return l
}

func (lex *Lexer) ReadRune() (rune, int, error) {
	lex.pos++

	current, size, err := lex.in.ReadRune()

	if current == '\n' {
		lex.line++
		lex.pos = 1
		return ' ', 1, nil
	}

	return current, size, err
}

// GetToken Get the next Token from the stream
func (lex *Lexer) GetToken() (*Token, error) {
	tok := new(Token)
	tok.Code = EOF

	if lex.eof {
		return tok, nil
	}

	var err error

	// Eat whitespace at beginning of buffer
	if unicode.IsSpace(lex.currentRune) {
		if lex.currentRune == '\n' {
			lex.line++
			lex.pos = 1
		}

		for lex.currentRune, _, err = lex.ReadRune(); unicode.IsSpace(lex.currentRune) && err == nil; {
			if lex.currentRune == '\n' {
				lex.line++
				lex.pos = 1
			}

			lex.currentRune, _, err = lex.ReadRune()
		}
	}

	// Just white space left ?
	if err != nil && err.Error() == "EOF" {
		return tok, nil
	}

	// Otherwise an error reading the next rune when slurping whitespace
	if err != nil {
		return nil, err
	}

	// Is alpha
	if unicode.IsLetter(lex.currentRune) {
		tok.Code = ATOM
		tok.Line = lex.line
		tok.Position = lex.pos

		var str = string(lex.currentRune)

		// Slurp letters and numbers up to something else
		lex.currentRune, _, err = lex.ReadRune()
		for unicode.IsLetter(lex.currentRune) && err == nil {
			str = str + string(lex.currentRune)
			lex.currentRune, _, err = lex.ReadRune()
		}

		if err != nil && err.Error() == "EOF" {
			lex.eof = true
		} else if err != nil {
			// eek non eof error
			return nil, err
		}

		tok.Code = StringToTokenType(str)

		if tok.Code == ATOM {
			tok.Value = str
		}

		return tok, nil
	}

	// Single char tokens, set the value if the token we created at the top
	tok.Line = lex.line
	tok.Position = lex.pos

	// plus
	switch lex.currentRune {
	case '\'':
		tok.Code = SHORTQUOTE
	case '(':
		tok.Code = OPENPARENS
	case ')':
		tok.Code = CLOSEPARENS
	default:
		return nil, fmt.Errorf("unrecognised token: %q", lex.currentRune)
	}

	// Move to the next rune
	lex.currentRune, _, err = lex.ReadRune()

	if err != nil && err.Error() == "EOF" {
		lex.eof = true
	} else if err != nil {
		// Error reading next rune, return an error instead of the last token
		return nil, err
	}

	return tok, nil
}
