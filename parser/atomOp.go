package parser

import "fmt"

// A List

type AtomOp struct {
	BaseNode
	entries []Node
}

func NewAtomOp(line int, position int) *AtomOp {
	return &AtomOp{BaseNode: BaseNode{line, position}}
}

func (op *AtomOp) String() string {
	return "Atom Operator"
}

func (op *AtomOp) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *AtomOp) Line() int {
	return op.BaseNode.Line
}

func (op *AtomOp) Position() int {
	return op.BaseNode.Position
}

func (op *AtomOp) Children() []Node {
	return op.entries
}

func (op *AtomOp) SyntaxCheck() error {
	// Only one argument for quote
	if len(op.entries) != 1 {
		return fmt.Errorf("atom operator requires only 1 argument, line %d, position %d", op.Line(), op.Position())
	}
	return nil
}

func (op *AtomOp) Evaluate() (Node, error) {
	return nil, nil
}
