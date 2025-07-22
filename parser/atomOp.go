package parser

// A List

type AtomOp struct {
	BaseNode
	entries []Node
}

func NewAtomOp(line int, position int) *AtomOp {
	return &AtomOp{BaseNode: BaseNode{line, position}}
}

func (op *AtomOp) String() string {
	return "Atom Operator"
}

func (op *AtomOp) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *AtomOp) Children() []Node {
	return op.entries
}

func (op *AtomOp) Evaluate() (Node, error) {
	return nil, nil
}
