package parser

// A List

type CdrOp struct {
	BaseNode
	entries []Node
}

func NewCdrOp(line int, position int) *CdrOp {
	return &CdrOp{BaseNode: BaseNode{line, position}}
}

func (op *CdrOp) String() string {
	return "Cdr Operator"
}

func (op *CdrOp) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *CdrOp) Children() []Node {
	return op.entries
}

func (op *CdrOp) Evaluate() (Node, error) {
	return nil, nil
}
