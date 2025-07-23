package parser

import "fmt"

// A List

type QuoteOp struct {
	BaseNode
	entries []Node
}

func NewQuoteOp(line int, position int) *QuoteOp {
	return &QuoteOp{BaseNode: BaseNode{line, position}}
}

func (op *QuoteOp) String() string {
	return "Quote Operator"
}

func (op *QuoteOp) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *QuoteOp) Line() int {
	return op.BaseNode.Line
}

func (op *QuoteOp) Position() int {
	return op.BaseNode.Position
}

func (op *QuoteOp) Children() []Node {
	return op.entries
}

func (op *QuoteOp) SyntaxCheck() error {
	// Only one argument for quote
	if len(op.entries) != 1 {
		return fmt.Errorf("quote operator requires only 1 argument, line %d, position %d", op.Line, op.Position)
	}
	return nil
}

func (op *QuoteOp) Evaluate() (Node, error) {
	return nil, nil
}
