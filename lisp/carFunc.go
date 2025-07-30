package lisp

import "fmt"

// A List

type CarFunc struct {
	BaseNode
}

func NewCarOp(line int, position int) *CarFunc {
	return &CarFunc{BaseNode: BaseNode{line, position}}
}

func (op *CarFunc) String() string {
	return "Car Operator"
}

// Interface Node
func (op *CarFunc) NodeType() string {
	return "Car Function"
}

func (op *CarFunc) Line() int {
	return op.BaseNode.Line
}

func (op *CarFunc) Position() int {
	return op.BaseNode.Position
}

func (op *CarFunc) Run(args []Node, env *Env) (Node, error) {
	// Only one argument for quote
	if len(args) != 1 {
		return nil, fmt.Errorf("car operator requires only 1 argument")
	}

	retNode, err := EvaluateNode(args[0], env, false)

	if err != nil {
		return nil, err
	}

	if listNode, ok := retNode.(*List); ok {
		// Empty list, returns as an empty list
		if len(listNode.Children()) == 0 {
			return NewList(0, 0), nil
		}

		return listNode.Children()[0], nil
	}

	return nil, fmt.Errorf("car operator requires a list as its argument")
}
