package lexer

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// Lexer Simple lex
type Lexer struct {
	in          *bufio.Reader
	currentRune rune
	eof         bool
}

// NewLexer Create a new Lexer
func NewLexer(input *bufio.Reader) *Lexer {
	l := new(Lexer)

	l.in = input
	l.currentRune = ' '
	l.eof = false

	return l
}

// GetToken Get the next Token from the stream
func (lex *Lexer) GetToken() (*Token, error) {
	tok := new(Token)
	tok.Code = TokEOF

	if lex.eof {
		return tok, nil
	}

	var err error

	// Eat whitespace at beginning of buffer
	if unicode.IsSpace(lex.currentRune) {
		for lex.currentRune, _, err = lex.in.ReadRune(); unicode.IsSpace(lex.currentRune) && err == nil; {
			lex.currentRune, _, err = lex.in.ReadRune()
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
		tok.Code = TokSTRING
		var str = string(lex.currentRune)

		// Slurp letters and numbers upto something else
		for lex.currentRune, _, err = lex.in.ReadRune(); (unicode.IsLetter(lex.currentRune) || unicode.IsNumber(lex.currentRune)) && err == nil; {
			str = str + string(lex.currentRune)
			lex.currentRune, _, err = lex.in.ReadRune()
		}

		if err != nil && err.Error() == "EOF" {
			lex.eof = true
		} else if err != nil {
			// eek non eof error
			return nil, err
		}

		tok.StringValue = str
		return tok, nil
	}

	// number
	if unicode.IsNumber(lex.currentRune) {
		var isFloat = false
		var str = string(lex.currentRune)

		for lex.currentRune, _, err = lex.in.ReadRune(); (unicode.IsNumber(lex.currentRune) == true || lex.currentRune == '.') && err == nil; {
			if lex.currentRune == '.' {
				if isFloat {
					// Hmm we've seen a . already error !
					return nil, errors.New("Double '.' in number")
				}

				isFloat = true
			}

			str = str + string(lex.currentRune)
			lex.currentRune, _, err = lex.in.ReadRune()

		}

		if err != nil && err.Error() == "EOF" {
			lex.eof = true
		} else if err != nil {
			// eek non eof error
			return nil, err
		}

		if isFloat {
			tok.Code = TokFLOAT
			tok.FloatValue, _ = strconv.ParseFloat(str, 64)
		} else {
			tok.Code = TokINT
			tok.IntValue, _ = strconv.ParseInt(str, 10, 64)
		}

		return tok, nil
	}

	// Single char tokens, set the value if the token we created at the top

	// plus
	switch lex.currentRune {
	case '+':
		tok.Code = TokPLUS
	case '-':
		tok.Code = TokMINUS
	case '*':
		tok.Code = TokMULTIPLY
	case '/':
		tok.Code = TokDIVIDE
	case '{':
		tok.Code = TokOPENBRACE
	case '}':
		tok.Code = TokCLOSEBRACE
	case '(':
		tok.Code = TokOPENPARENS
	case ')':
		tok.Code = TokCLOSEPARENS
	default:
		return nil, fmt.Errorf("Unrecoginsed token: %v", lex.currentRune)
	}

	// Move to the next rune
	lex.currentRune, _, err = lex.in.ReadRune()

	if err != nil && err.Error() == "EOF" {
		lex.eof = true
	} else if err != nil {
		// Error reading next rune, return an error instead of the last token
		return nil, err
	}

	return tok, nil
}
