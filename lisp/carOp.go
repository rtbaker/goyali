package lisp

import "fmt"

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
func (op *CarOp) QuotedValue() Node {
	return NewAtom("car", op.Line(), op.Position())
}

func (op *CarOp) Line() int {
	return op.BaseNode.Line
}

func (op *CarOp) Position() int {
	return op.BaseNode.Position
}

func (op *CarOp) Children() []Node {
	return op.entries
}

func (op *CarOp) SyntaxCheck() error {
	// Only one argument for car
	if len(op.entries) != 1 {
		return fmt.Errorf("car operator requires only 1 argument, line %d, position %d", op.Line(), op.Position())
	}
	return nil
}

func (op *CarOp) Evaluate() (Node, error) {
	return nil, nil
}
