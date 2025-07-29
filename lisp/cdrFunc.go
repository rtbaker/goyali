package lisp

import "fmt"

// A List

type CdrFunc struct {
	BaseNode
	entries []Node
}

func NewCdrOp(line int, position int) *CdrFunc {
	return &CdrFunc{BaseNode: BaseNode{line, position}}
}

func (op *CdrFunc) String() string {
	return "Cdr Operator"
}

func (op *CdrFunc) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *CdrFunc) QuotedValue() Node {
	return NewAtom("cdr", op.Line(), op.Position())
}

func (op *CdrFunc) Line() int {
	return op.BaseNode.Line
}

func (op *CdrFunc) Position() int {
	return op.BaseNode.Position
}

func (op *CdrFunc) Children() []Node {
	return op.entries
}

func (op *CdrFunc) SyntaxCheck() error {
	// Only one argument for cdr
	if len(op.entries) != 1 {
		return fmt.Errorf("cdr operator requires only 1 argument, line %d, position %d", op.Line(), op.Position())
	}
	return nil
}

func (op *CdrFunc) Evaluate() (Node, error) {
	return nil, nil
}
