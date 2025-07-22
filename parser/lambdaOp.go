package parser

// A List

type LambdaOp struct {
	BaseNode
	entries []Node
}

func NewLambdaOp(line int, position int) *LambdaOp {
	return &LambdaOp{BaseNode: BaseNode{line, position}}
}

func (op *LambdaOp) String() string {
	return "Lambda Operator"
}

func (op *LambdaOp) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *LambdaOp) Children() []Node {
	return op.entries
}

func (op *LambdaOp) Evaluate() (Node, error) {
	return nil, nil
}
