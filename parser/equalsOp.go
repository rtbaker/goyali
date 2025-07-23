package parser

import "fmt"

// A List

type EqualsOp struct {
	BaseNode
	entries []Node
}

func NewEqualsOp(line int, position int) *EqualsOp {
	return &EqualsOp{BaseNode: BaseNode{line, position}}
}

func (op *EqualsOp) String() string {
	return "Equals Operator"
}

func (op *EqualsOp) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *EqualsOp) Children() []Node {
	return op.entries
}

func (op *EqualsOp) SyntaxCheck() error {
	// Only 2 arguments for eq
	if len(op.entries) != 2 {
		return fmt.Errorf("equals operator requires 2 arguments, line %d, position %d", op.Line, op.Position)
	}
	return nil
}

func (op *EqualsOp) Evaluate() (Node, error) {
	return nil, nil
}
