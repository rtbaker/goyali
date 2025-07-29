package lisp

import "fmt"

// A List

type LabelFunc struct {
	BaseNode
	entries []Node
}

func NewLabelOp(line int, position int) *LabelFunc {
	return &LabelFunc{BaseNode: BaseNode{line, position}}
}

func (op *LabelFunc) String() string {
	return "Label Operator"
}

func (op *LabelFunc) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *LabelFunc) NodeType() string {
	return "Label Function"
}

func (op *LabelFunc) QuotedValue() Node {
	return NewAtom("label", op.Line(), op.Position())
}

func (op *LabelFunc) Line() int {
	return op.BaseNode.Line
}

func (op *LabelFunc) Position() int {
	return op.BaseNode.Position
}

func (op *LabelFunc) Children() []Node {
	return op.entries
}

func (op *LabelFunc) SyntaxCheck() error {
	if len(op.entries) != 2 {
		return fmt.Errorf("label op requires 2 arguments, line %d position %d", op.Line(), op.Position())
	}

	labelNode := op.entries[0]

	if _, ok := labelNode.(*Atom); !ok {
		return fmt.Errorf("label op first arg must be an atom, line %d position %d", op.Line(), op.Position())
	}

	return nil
}

func (op *LabelFunc) Run(args []Node) (Node, error) {
	return nil, nil
}
