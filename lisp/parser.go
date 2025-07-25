package lisp

import (
	"fmt"

	"github.com/rtbaker/goyali/lexer"
)

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

func (parser *Parser) Parse() (Node, error) {
	top := &Program{}

	var err error
	parser.lookahead, err = parser.lexer.GetToken()

	if err != nil {
		return nil, err
	}

	for expr, err := parser.getExpression(); expr != nil; expr, err = parser.getExpression() {
		if err != nil {
			return nil, err
		}

		top.expressions = append(top.expressions, expr)
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
		quote := NewQuoteOp(parser.lookahead.Line, parser.lookahead.Position)

		err := parser.match(lexer.SHORTQUOTE)
		if err != nil {
			return nil, err
		}

		expression, err := parser.getExpression()
		if err != nil {
			return nil, err
		}

		quote.AppendNode(expression)

		return quote, nil
	}

	// Shouldn't get here
	return nil, fmt.Errorf(
		"syntax error, mis-formed expression, line: %d, character: %d",
		parser.lookahead.Line, parser.lookahead.Position)
}

// A list contains atoms or other lists
// Some of the lists are special (i.e. start with language keyword)
func (parser *Parser) getList() (Node, error) {
	// Start
	line := parser.lookahead.Line
	pos := parser.lookahead.Position

	err := parser.match(lexer.OPENPARENS)
	if err != nil {
		return nil, fmt.Errorf("list start error: %s", err)
	}

	list, err := parser.getOperator(line, pos)
	if err != nil {
		return nil, fmt.Errorf("operator match error: %s", err)
	}

	// Just an ordinary list?
	if list == nil {
		list = NewList(line, pos)
	}

	// Keep adding children/entries until the list is closed
	for parser.lookahead.Code != lexer.CLOSEPARENS {
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

// If the list starts with a known operator return it's
// type or nil. Assume the initial parens has been consumed
// so we are looking for a know op type
func (parser *Parser) getOperator(line int, position int) (ListNode, error) {
	var node ListNode = nil
	var err error

	switch parser.lookahead.Code {
	case lexer.QUOTE:
		node = NewQuoteOp(line, position)
		err = parser.match(lexer.QUOTE)
	case lexer.ATOMOPERATOR:
		node = NewAtomOp(line, position)
		err = parser.match(lexer.ATOMOPERATOR)
	case lexer.EQUALS:
		node = NewEqualsOp(line, position)
		err = parser.match(lexer.EQUALS)
	case lexer.CAR:
		node = NewCarOp(line, position)
		err = parser.match(lexer.CAR)
	case lexer.CDR:
		node = NewCdrOp(line, position)
		err = parser.match(lexer.CDR)
	case lexer.CONS:
		node = NewConsOp(line, position)
		err = parser.match(lexer.CONS)
	case lexer.COND:
		node = NewCondOp(line, position)
		err = parser.match(lexer.COND)
	case lexer.LAMBDA:
		node = NewLambdaOp(line, position)
		err = parser.match(lexer.LAMBDA)
	case lexer.LABEL:
		node = NewLabelOp(line, position)
		err = parser.match(lexer.LABEL)
	case lexer.DEFUN:
		node = NewDefunOp(line, position)
		err = parser.match(lexer.DEFUN)
	}

	// didn't match a know operator
	return node, err
}

// Checks that we have what we are expecting and moves us to the next
// token if so.
func (parser *Parser) match(nodeType lexer.TokenCode) error {
	if parser.lookahead.Code == nodeType {
		var err error
		parser.lookahead, err = parser.lexer.GetToken()
		return err
	}

	return fmt.Errorf("expected token code: %d, got %d at line %d, position %d",
		nodeType, parser.lookahead.Code, parser.lookahead.Line, parser.lookahead.Position)
}
