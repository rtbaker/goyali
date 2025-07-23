package parser

// A List

type LabelOp struct {
	BaseNode
	entries []Node
}

func NewLabelOp(line int, position int) *LabelOp {
	return &LabelOp{BaseNode: BaseNode{line, position}}
}

func (op *LabelOp) String() string {
	return "Label Operator"
}

func (op *LabelOp) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *LabelOp) Children() []Node {
	return op.entries
}

func (op *LabelOp) SyntaxCheck() error {
	return nil
}

func (op *LabelOp) Evaluate() (Node, error) {
	return nil, nil
}
