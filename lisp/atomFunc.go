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
func (op *AtomFunc) NodeType() string {
	return "Atom Function"
}

func (op *AtomFunc) QuotedValue() Node {
	return NewAtom("atom", op.Line(), op.Position())
}

func (op *AtomFunc) Line() int {
	return op.BaseNode.Line
}

func (op *AtomFunc) Position() int {
	return op.BaseNode.Position
}

func (op *AtomFunc) Run(args []Node, env *Env) (Node, error) {
	// Only one argument for quote
	if len(args) != 1 {
		return nil, fmt.Errorf("atom operator requires only 1 argument")
	}

	retNode, err := EvaluateNode(args[0], env, false)

	if err != nil {
		return nil, err
	}

	if NodeIsAtom(retNode) {
		return Truth(), nil
	}

	if listVal, ok := retNode.(*List); ok {
		if listVal.isEmptyList() {
			return Truth(), nil
		}
	}

	return Falsity(), nil
}
