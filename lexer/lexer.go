package lexer

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"
)

// Lexer Simple lex
type Lexer struct {
	in          io.Reader
	scanner     *bufio.Scanner
	currentRune rune
	eof         bool
	pos         int
	line        int
}

// NewLexer Create a new Lexer
func NewLexer(input io.Reader) *Lexer {
	l := new(Lexer)

	l.in = input
	l.scanner = bufio.NewScanner(l.in)
	l.scanner.Split(bufio.ScanRunes)

	l.eof = false
	l.pos = 0
	l.line = 1
	l.currentRune = ' '

	return l
}

func (lex *Lexer) ResetLineNos() {
	lex.pos = 0
	lex.line = 1
}

func (lex *Lexer) ReadNextRune() (rune, error) {
	lex.pos++

	ok := lex.scanner.Scan()

	if !ok {
		// either end of file or an error
		err := lex.scanner.Err()

		if err != nil {
			return utf8.RuneError, err
		}

		return utf8.RuneError, io.EOF
	}

	b := lex.scanner.Bytes()
	current, _ := utf8.DecodeRune(b)

	if current == '\n' {
		lex.line++
		lex.pos = 1
	}

	return current, nil
}

// GetToken Get the next Token from the stream
func (lex *Lexer) GetToken() (*Token, error) {
	// default end of file token
	tok := new(Token)
	tok.Code = EOF
	tok.Line = lex.line
	tok.Position = lex.pos

	if lex.eof {
		return tok, nil
	}

	var err error

	// Eat whitespace at beginning of buffer
	if unicode.IsSpace(lex.currentRune) {
		if lex.currentRune == '\n' {
			lex.line++
			lex.pos = 0
		}

		for lex.currentRune, err = lex.ReadNextRune(); unicode.IsSpace(lex.currentRune) && err == nil; {
			if lex.currentRune == '\n' {
				lex.line++
				lex.pos = 0
			}

			lex.currentRune, err = lex.ReadNextRune()
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

	// Is alpha (and '.')
	if unicode.IsLetter(lex.currentRune) || lex.currentRune == '.' {
		tok.Code = ATOM
		tok.Line = lex.line
		tok.Position = lex.pos

		var str = string(lex.currentRune)

		// Slurp letters and numbers up to something else
		lex.currentRune, err = lex.ReadNextRune()
		for unicode.IsLetter(lex.currentRune) || lex.currentRune == '.' && err == nil {
			str = str + string(lex.currentRune)
			lex.currentRune, err = lex.ReadNextRune()
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
	case ';':
		err := lex.slurpToEOL()
		if err != nil {
			return nil, err
		}
		return lex.GetToken()
	default:
		return nil, fmt.Errorf("unrecognised token: %q", lex.currentRune)
	}

	// Move to the next rune
	lex.currentRune, err = lex.ReadNextRune()

	if err != nil && err.Error() == "EOF" {
		lex.eof = true
	} else if err != nil {
		// Error reading next rune, return an error instead of the last token
		return nil, err
	}

	return tok, nil
}

// remove all chars to end of line
func (lex *Lexer) slurpToEOL() error {
	var err error

	for lex.currentRune, err = lex.ReadNextRune(); lex.currentRune != '\n' && err == nil; {
		lex.currentRune, err = lex.ReadNextRune()
	}

	if err != nil && err.Error() == "EOF" {
		lex.eof = true
		return nil
	}

	// Otherwise an error reading the next rune when slurping whitespace
	if err != nil {
		return err
	}

	return nil
}
