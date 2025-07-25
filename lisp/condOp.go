package lisp

import "fmt"

// A List

type CondOp struct {
	BaseNode
	entries []Node
}

func NewCondOp(line int, position int) *CondOp {
	return &CondOp{BaseNode: BaseNode{line, position}}
}

func (op *CondOp) String() string {
	return "Cond Operator"
}

func (op *CondOp) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *CondOp) QuotedValue() Node {
	return NewAtom("cond", op.Line(), op.Position())
}

func (op *CondOp) Line() int {
	return op.BaseNode.Line
}

func (op *CondOp) Position() int {
	return op.BaseNode.Position
}

func (op *CondOp) Children() []Node {
	return op.entries
}

func (op *CondOp) SyntaxCheck() error {
	// Must have at least 1 entry?
	if len(op.entries) == 0 {
		return fmt.Errorf("cond operator requires at least 1 argument, line %d position %d", op.Line(), op.Position())
	}

	// each entry is (<test expression> <return expression>)
	for _, n := range op.Children() {
		if len(n.Children()) != 2 {
			return fmt.Errorf("cond entry requires 2 expressions, line %d position %d", n.Line(), n.Position())
		}
	}
	return nil
}

func (op *CondOp) Evaluate() (Node, error) {
	return nil, nil
}
