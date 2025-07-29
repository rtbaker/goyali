package lisp

import "fmt"

// A List

type CondFunc struct {
	BaseNode
	entries []Node
}

func NewCondOp(line int, position int) *CondFunc {
	return &CondFunc{BaseNode: BaseNode{line, position}}
}

func (op *CondFunc) String() string {
	return "Cond Operator"
}

func (op *CondFunc) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *CondFunc) NodeType() string {
	return "Cond Function"
}

func (op *CondFunc) QuotedValue() Node {
	return NewAtom("cond", op.Line(), op.Position())
}

func (op *CondFunc) Line() int {
	return op.BaseNode.Line
}

func (op *CondFunc) Position() int {
	return op.BaseNode.Position
}

func (op *CondFunc) Children() []Node {
	return op.entries
}

func (op *CondFunc) SyntaxCheck() error {
	// Must have at least 1 entry?
	if len(op.entries) == 0 {
		return fmt.Errorf("cond operator requires at least 1 argument, line %d position %d", op.Line(), op.Position())
	}

	// each entry is (<test expression> <return expression>)
	//	for _, n := range op.Children() {
	//		if len(n.Children()) != 2 {
	//	return fmt.Errorf("cond entry requires 2 expressions, line %d position %d", n.Line(), n.Position())
	//	}
	//	}
	return nil
}

func (op *CondFunc) Run(args []Node) (Node, error) {
	return nil, nil
}
