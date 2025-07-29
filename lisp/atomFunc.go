package lisp

import "fmt"

// A List

type AtomFunc struct {
	BaseNode
	entries []Node
}

func NewAtomOp(line int, position int) *AtomFunc {
	return &AtomFunc{BaseNode: BaseNode{line, position}}
}

func (op *AtomFunc) String() string {
	return "Atom Operator"
}

func (op *AtomFunc) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *AtomFunc) QuotedValue() Node {
	return NewAtom("atom", op.Line(), op.Position())
}

func (op *AtomFunc) Line() int {
	return op.BaseNode.Line
}

func (op *AtomFunc) Position() int {
	return op.BaseNode.Position
}

func (op *AtomFunc) Children() []Node {
	return op.entries
}

func (op *AtomFunc) SyntaxCheck() error {
	// Only one argument for quote
	if len(op.entries) != 1 {
		return fmt.Errorf("atom operator requires only 1 argument, line %d, position %d", op.Line(), op.Position())
	}
	return nil
}

func (op *AtomFunc) Evaluate() (Node, error) {
	return nil, nil
}
