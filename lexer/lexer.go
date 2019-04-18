package lexer

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

type TokenCode int

const (
	Tok_eof        TokenCode = 1
	Tok_int        TokenCode = 2
	Tok_float      TokenCode = 3
	Tok_plus       TokenCode = 4
	Tok_minus      TokenCode = 5
	Tok_multiply   TokenCode = 6
	Tok_divide     TokenCode = 7
	Tok_openbrace  TokenCode = 8
	Tok_closebrace TokenCode = 9
	Tok_string     TokenCode = 10
)

type Token struct {
	Code        TokenCode
	IntValue    int64
	FloatValue  float64
	StringValue string
}

type lexer struct {
	in          *bufio.Reader
	currentRune rune
}

func NewLexer(input *bufio.Reader) *lexer {
	l := new(lexer)

	l.in = input
	l.currentRune = ' '

	return l
}

func (lex *lexer) GetToken() (*Token, error) {
	tok := new(Token)
	tok.Code = Tok_eof

	var err error

	// Eat whitespace
	if unicode.IsSpace(lex.currentRune) {
		for lex.currentRune, _, err = lex.in.ReadRune(); unicode.IsSpace(lex.currentRune) && err != nil; {
			lex.currentRune, _, err = lex.in.ReadRune()
		}
	}

	if err != nil {
		return nil, err
	}

	// Is alpha
	if unicode.IsLetter(lex.currentRune) {
		tok.Code = Tok_string
		var str string = string(lex.currentRune)

		// Slurp letters and numbers upto something else
		for lex.currentRune, _, err = lex.in.ReadRune(); unicode.IsLetter(lex.currentRune) || unicode.IsNumber(lex.currentRune); {
			if err != nil {
				return nil, err
			}

			str = str + string(lex.currentRune)
			lex.currentRune, _, err = lex.in.ReadRune()
		}

		tok.StringValue = str
		return tok, nil
	}

	// number
	if unicode.IsNumber(lex.currentRune) {
		var isFloat bool = false
		var str string = string(lex.currentRune)

		for lex.currentRune, _, err = lex.in.ReadRune(); unicode.IsNumber(lex.currentRune) == true || lex.currentRune == '.'; {
			if err != nil {
				return nil, err
			}

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

		if isFloat {
			tok.Code = Tok_float
			tok.FloatValue, _ = strconv.ParseFloat(str, 64)
		} else {
			tok.Code = Tok_int
			tok.IntValue, _ = strconv.ParseInt(str, 10, 64)
		}

		return tok, nil
	}

	// plus
	if lex.currentRune == '+' {
		tok.Code = Tok_plus

		// advance stream
		lex.currentRune, _, err = lex.in.ReadRune()

		if err != nil {
			return nil, err
		}
	}

	// plus
	if lex.currentRune == '-' {
		tok.Code = Tok_minus

		// advance stream
		lex.currentRune, _, err = lex.in.ReadRune()

		if err != nil {
			return nil, err
		}
	}

	// plus
	if lex.currentRune == '*' {
		tok.Code = Tok_multiply

		// advance stream
		lex.currentRune, _, err = lex.in.ReadRune()

		if err != nil {
			return nil, err
		}
	}

	// plus
	if lex.currentRune == '/' {
		tok.Code = Tok_divide

		// advance stream
		lex.currentRune, _, err = lex.in.ReadRune()

		if err != nil {
			return nil, err
		}
	}

	return tok, nil
}

func (tok *Token) PrintToken() {
	switch tok.Code {
	case Tok_eof:
		fmt.Println("End of file")
	case Tok_int:
		fmt.Printf("Integer: %d\n", tok.IntValue)
	case Tok_float:
		fmt.Printf("Float: %f\n", tok.FloatValue)
	case Tok_plus:
		fmt.Println("Plus")
	case Tok_minus:
		fmt.Println("Minus")
	case Tok_multiply:
		fmt.Println("Multiply")
	case Tok_divide:
		fmt.Println("Divide")
	case Tok_openbrace:
		fmt.Println("Open brace")
	case Tok_closebrace:
		fmt.Println("Close brace")
	case Tok_string:
		fmt.Printf("String: %s\n", tok.StringValue)
	}
}
