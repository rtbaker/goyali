package parser

// A List

type ConsOp struct {
	BaseNode
	entries []Node
}

func NewConsOp(line int, position int) *ConsOp {
	return &ConsOp{BaseNode: BaseNode{line, position}}
}

func (op *ConsOp) String() string {
	return "Cons Operator"
}

func (op *ConsOp) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *ConsOp) Children() []Node {
	return op.entries
}

func (op *ConsOp) SyntaxCheck() error {
	return nil
}

func (op *ConsOp) Evaluate() (Node, error) {
	return nil, nil
}
