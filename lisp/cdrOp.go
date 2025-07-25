package lisp

import "fmt"

// A List

type CdrOp struct {
	BaseNode
	entries []Node
}

func NewCdrOp(line int, position int) *CdrOp {
	return &CdrOp{BaseNode: BaseNode{line, position}}
}

func (op *CdrOp) String() string {
	return "Cdr Operator"
}

func (op *CdrOp) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *CdrOp) Line() int {
	return op.BaseNode.Line
}

func (op *CdrOp) Position() int {
	return op.BaseNode.Position
}

func (op *CdrOp) Children() []Node {
	return op.entries
}

func (op *CdrOp) SyntaxCheck() error {
	// Only one argument for cdr
	if len(op.entries) != 1 {
		return fmt.Errorf("cdr operator requires only 1 argument, line %d, position %d", op.Line(), op.Position())
	}
	return nil
}

func (op *CdrOp) Evaluate() (Node, error) {
	return nil, nil
}
