package parser

// A List

type CondOp struct {
	BaseNode
	entries []Node
}

func NewCondOp(line int, position int) *CondOp {
	return &CondOp{BaseNode: BaseNode{line, position}}
}

func (op *CondOp) String() string {
	return "Cond Operator"
}

func (op *CondOp) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *CondOp) Children() []Node {
	return op.entries
}

func (op *CondOp) Evaluate() (Node, error) {
	return nil, nil
}
