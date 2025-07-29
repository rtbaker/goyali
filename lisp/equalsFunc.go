package lisp

import "fmt"

// A List

type EqualsFunc struct {
	BaseNode
	entries []Node
}

func NewEqualsOp(line int, position int) *EqualsFunc {
	return &EqualsFunc{BaseNode: BaseNode{line, position}}
}

func (op *EqualsFunc) String() string {
	return "Equals Operator"
}

func (op *EqualsFunc) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *EqualsFunc) QuotedValue() Node {
	return NewAtom("eq", op.Line(), op.Position())
}

func (op *EqualsFunc) Line() int {
	return op.BaseNode.Line
}

func (op *EqualsFunc) Position() int {
	return op.BaseNode.Position
}

func (op *EqualsFunc) Children() []Node {
	return op.entries
}

func (op *EqualsFunc) SyntaxCheck() error {
	// Only 2 arguments for eq
	if len(op.entries) != 2 {
		return fmt.Errorf("equals operator requires 2 arguments, line %d, position %d", op.Line(), op.Position())
	}
	return nil
}

func (op *EqualsFunc) Evaluate() (Node, error) {
	return nil, nil
}
