package parser

import (
	"fmt"

	"github.com/rtbaker/goyali/lexer"
)

type Parser struct {
	lexer     lexer.Lexer
	lookahead *lexer.Token
}

func NewParser(lexer lexer.Lexer) *Parser {
	parser := &Parser{
		lexer: lexer,
	}

	return parser
}

func (parser *Parser) Parse() (Node, error) {
	top := &Program{}

	var err error
	parser.lookahead, err = parser.lexer.GetToken()

	if err != nil {
		return nil, err
	}

	for expr, err := parser.getExpression(); expr != nil; {
		if err != nil {
			return nil, err
		}

		top.expressions = append(top.expressions, expr)
	}
	return top, nil
}

func (parser *Parser) getExpression() (Node, error) {
	if parser.lookahead.Code == lexer.ATOM {
		atom := NewAtom(parser.lookahead.Value)

		// Consume token
		var err error
		parser.lookahead, err = parser.lexer.GetToken()

		if err != nil {
			return nil, err
		}

		return atom, nil
	}

	if parser.lookahead.Code == lexer.OPENPARENS {
		list, err := parser.getList()

		if err != nil {
			return nil, err
		}

		return list, nil
	}

	// Shouldn't get here
	return nil, fmt.Errorf(
		"syntax error, mis-formed expression, line: %d, character: %d",
		parser.lookahead.Line, parser.lookahead.Position)
}

func (parser *Parser) getList() (Node, error) {
	return nil, nil
}
