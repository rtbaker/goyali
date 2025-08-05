package lisp

import "fmt"

// A List

type ConsOp struct {
	BaseNode
}

func NewConsOp(line int, position int) *ConsOp {
	return &ConsOp{BaseNode: BaseNode{line, position}}
}

func (op *ConsOp) String() string {
	return "Cons Operator"
}

// Interface Node
func (op *ConsOp) NodeType() string {
	return "Cons Function"
}

func (op *ConsOp) Line() int {
	return op.BaseNode.Line
}

func (op *ConsOp) Position() int {
	return op.BaseNode.Position
}

func (op *ConsOp) Run(args []Node, env *Env) (Node, error) {
	// Only one argument for quote
	if len(args) != 2 {
		return nil, fmt.Errorf("cons operator requires 2 arguments")
	}

	retNode1, err := EvaluateNode(args[0], env, false)

	if err != nil {
		return nil, err
	}

	retNode2, err := EvaluateNode(args[1], env, false)

	if err != nil {
		return nil, err
	}

	// arg 2 must be a list
	var listArg ListNode
	var ok bool

	if listArg, ok = retNode2.(*List); !ok {
		return nil, fmt.Errorf("cons operator requires a list for the second argument")
	}

	retList := NewList(0, 0)
	retList.AppendNode(retNode1)
	retList.AppendNodes(listArg.Children())

	return retList, nil
}
