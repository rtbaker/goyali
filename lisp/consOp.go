package lisp

import "fmt"

// A List

type ConsOp struct {
	BaseNode
	entries []Node
}

func NewConsOp(line int, position int) *ConsOp {
	return &ConsOp{BaseNode: BaseNode{line, position}}
}

func (op *ConsOp) String() string {
	return "Cons Operator"
}

func (op *ConsOp) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *ConsOp) QuotedValue() Node {
	return NewAtom("cons", op.Line(), op.Position())
}

func (op *ConsOp) Line() int {
	return op.BaseNode.Line
}

func (op *ConsOp) Position() int {
	return op.BaseNode.Position
}

func (op *ConsOp) Children() []Node {
	return op.entries
}

func (op *ConsOp) SyntaxCheck() error {
	// Only 2 arguments for cons
	if len(op.entries) != 2 {
		return fmt.Errorf("cons operator requires 2 arguments, line %d, position %d", op.Line(), op.Position())
	}
	return nil
}

func (op *ConsOp) Evaluate() (Node, error) {
	return nil, nil
}
