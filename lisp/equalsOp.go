package lisp

import "fmt"

// A List

type EqualsOp struct {
	BaseNode
}

func NewEqualsOp(line int, position int) *EqualsOp {
	return &EqualsOp{BaseNode: BaseNode{line, position}}
}

func (op *EqualsOp) String() string {
	return "Equals Operator"
}

// Interface Node
func (op *EqualsOp) NodeType() string {
	return "Equals Function"
}

func (op *EqualsOp) Line() int {
	return op.BaseNode.Line
}

func (op *EqualsOp) Position() int {
	return op.BaseNode.Position
}

func (op *EqualsOp) Run(args []Node, env *Env) (Node, error) {
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
