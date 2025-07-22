package parser

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
func (op *QuoteOp) Children() []Node {
	return op.entries
}

func (op *QuoteOp) Evaluate() (Node, error) {
	return nil, nil
}
