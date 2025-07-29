package lisp

import "fmt"

// A List

type ConsFunc struct {
	BaseNode
	entries []Node
}

func NewConsOp(line int, position int) *ConsFunc {
	return &ConsFunc{BaseNode: BaseNode{line, position}}
}

func (op *ConsFunc) String() string {
	return "Cons Operator"
}

func (op *ConsFunc) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *ConsFunc) QuotedValue() Node {
	return NewAtom("cons", op.Line(), op.Position())
}

func (op *ConsFunc) Line() int {
	return op.BaseNode.Line
}

func (op *ConsFunc) Position() int {
	return op.BaseNode.Position
}

func (op *ConsFunc) Children() []Node {
	return op.entries
}

func (op *ConsFunc) SyntaxCheck() error {
	// Only 2 arguments for cons
	if len(op.entries) != 2 {
		return fmt.Errorf("cons operator requires 2 arguments, line %d, position %d", op.Line(), op.Position())
	}
	return nil
}

func (op *ConsFunc) Evaluate() (Node, error) {
	return nil, nil
}
