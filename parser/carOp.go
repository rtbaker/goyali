package parser

// A List

type CarOp struct {
	BaseNode
	entries []Node
}

func NewCarOp(line int, position int) *CarOp {
	return &CarOp{BaseNode: BaseNode{line, position}}
}

func (op *CarOp) String() string {
	return "Car Operator"
}

func (op *CarOp) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *CarOp) Children() []Node {
	return op.entries
}

func (op *CarOp) Evaluate() (Node, error) {
	return nil, nil
}
