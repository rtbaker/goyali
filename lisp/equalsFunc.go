package lisp

import "fmt"

// A List

type EqualsFunc struct {
	BaseNode
	entries []Node
}

func NewEqualsOp(line int, position int) *EqualsFunc {
	return &EqualsFunc{BaseNode: BaseNode{line, position}}
}

func (op *EqualsFunc) String() string {
	return "Equals Operator"
}

func (op *EqualsFunc) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *EqualsFunc) NodeType() string {
	return "Equals Function"
}

func (op *EqualsFunc) QuotedValue() Node {
	return NewAtom("eq", op.Line(), op.Position())
}

func (op *EqualsFunc) Line() int {
	return op.BaseNode.Line
}

func (op *EqualsFunc) Position() int {
	return op.BaseNode.Position
}

func (op *EqualsFunc) Children() []Node {
	return op.entries
}

func (op *EqualsFunc) SyntaxCheck() error {
	// Only 2 arguments for eq
	if len(op.entries) != 2 {
		return fmt.Errorf("equals operator requires 2 arguments, line %d, position %d", op.Line(), op.Position())
	}
	return nil
}

func (op *EqualsFunc) Run(args []Node, env *Env) (Node, error) {
	// Only one argument for quote
	if len(args) != 2 {
		return nil, fmt.Errorf("equals operator requires 2 arguments")
	}

	retNode1, err := EvaluateNode(args[0], env, false)

	if err != nil {
		return nil, err
	}

	retNode2, err := EvaluateNode(args[1], env, false)

	if err != nil {
		return nil, err
	}

	// Both the same atom?
	if atom1, ok := retNode1.(*Atom); ok {
		if atom2, ok := retNode2.(*Atom); ok {
			if atom1.Name == atom2.Name {
				return Truth(), nil
			}
		}
	}

	// Both empty list?
	if list1, ok := retNode1.(*List); ok {
		if list2, ok := retNode2.(*List); ok {
			if list1.isEmptyList() && list2.isEmptyList() {
				return Truth(), nil
			}
		}
	}

	return Falsity(), nil
}
