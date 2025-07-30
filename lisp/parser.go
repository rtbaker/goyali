package lisp

import (
	"fmt"

	"github.com/rtbaker/goyali/lexer"
)

// Takes a stream of tokens from the lexer and returns lists and atoms

type Parser struct {
	lexer     *lexer.Lexer
	lookahead *lexer.Token
}

func NewParser(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer: lexer,
	}

	return parser
}

func (parser *Parser) ParseProgram() (*Program, error) {
	top := NewProgram()

	var err error
	parser.lookahead, err = parser.lexer.GetToken()

	if err != nil {
		return nil, err
	}

	var expr Node

	for expr, err = parser.getExpression(); expr != nil; expr, err = parser.getExpression() {
		if err != nil {
			return nil, err
		}

		top.AppendNode(expr)
	}

	// catch error in last iteration of loop
	if err != nil {
		return nil, err
	}

	return top, nil
}

func (parser *Parser) getExpression() (Node, error) {
	if parser.lookahead.Code == lexer.EOF {
		return nil, nil
	}

	if parser.lookahead.Code == lexer.ATOM {
		atom := NewAtom(parser.lookahead.Value, parser.lookahead.Line, parser.lookahead.Position)

		// Consume token
		err := parser.match(lexer.ATOM)
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

	if parser.lookahead.Code == lexer.SHORTQUOTE {
		line := parser.lookahead.Line
		pos := parser.lookahead.Position

		err := parser.match(lexer.SHORTQUOTE)
		if err != nil {
			return nil, err
		}

		list := NewList(line, pos)
		quoteAtom := NewAtom("quote", line, pos+1) // Line&position bug here, subsequent positions on this line are now wrong

		list.AppendNode(quoteAtom)

		expression, err := parser.getExpression()
		if err != nil {
			return nil, err
		}

		list.AppendNode(expression)

		return list, nil
	}

	// Shouldn't get here
	return nil, fmt.Errorf(
		"syntax error, mis-formed expression, line: %d, character: %d",
		parser.lookahead.Line, parser.lookahead.Position)
}

// A list contains atoms or other lists.Nod
func (parser *Parser) getList() (Node, error) {
	// Start
	line := parser.lookahead.Line
	pos := parser.lookahead.Position

	err := parser.match(lexer.OPENPARENS)
	if err != nil {
		return nil, fmt.Errorf("list start error: %s", err)
	}

	list := NewList(line, pos)

	// Keep adding children/entries until the list is closed
	for parser.lookahead.Code != lexer.CLOSEPARENS && parser.lookahead.Code != lexer.EOF {
		exp, err := parser.getExpression()
		if err != nil {
			return nil, err
		}
		list.AppendNode(exp)
	}

	// End
	err = parser.match(lexer.CLOSEPARENS)
	if err != nil {
		return nil, fmt.Errorf("list end error: %s", err)
	}

	return list, nil
}

// Checks that we have what we are expecting and moves us to the next
// token if so.
func (parser *Parser) match(nodeType lexer.TokenCode) error {
	if parser.lookahead.Code == nodeType {
		var err error
		parser.lookahead, err = parser.lexer.GetToken()
		return err
	}

	return fmt.Errorf("expected token code: %s, got %s at line %d, position %d",
		lexer.TokenCodeString(nodeType), lexer.TokenCodeString(parser.lookahead.Code), parser.lookahead.Line, parser.lookahead.Position)
}
