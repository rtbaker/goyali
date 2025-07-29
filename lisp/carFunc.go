package lisp

import "fmt"

// A List

type CarFunc struct {
	BaseNode
	entries []Node
}

func NewCarOp(line int, position int) *CarFunc {
	return &CarFunc{BaseNode: BaseNode{line, position}}
}

func (op *CarFunc) String() string {
	return "Car Operator"
}

func (op *CarFunc) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *CarFunc) NodeType() string {
	return "Car Function"
}

func (op *CarFunc) QuotedValue() Node {
	return NewAtom("car", op.Line(), op.Position())
}

func (op *CarFunc) Line() int {
	return op.BaseNode.Line
}

func (op *CarFunc) Position() int {
	return op.BaseNode.Position
}

func (op *CarFunc) Children() []Node {
	return op.entries
}

func (op *CarFunc) SyntaxCheck() error {
	// Only one argument for car
	if len(op.entries) != 1 {
		return fmt.Errorf("car operator requires only 1 argument, line %d, position %d", op.Line(), op.Position())
	}
	return nil
}

func (op *CarFunc) Run(args []Node) (Node, error) {
	return nil, nil
}
