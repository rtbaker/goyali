package lisp

import "fmt"

// A List

type ConsFunc struct {
	BaseNode
	entries []Node
}

func NewConsOp(line int, position int) *ConsFunc {
	return &ConsFunc{BaseNode: BaseNode{line, position}}
}

func (op *ConsFunc) String() string {
	return "Cons Operator"
}

func (op *ConsFunc) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *ConsFunc) NodeType() string {
	return "Cons Function"
}

func (op *ConsFunc) QuotedValue() Node {
	return NewAtom("cons", op.Line(), op.Position())
}

func (op *ConsFunc) Line() int {
	return op.BaseNode.Line
}

func (op *ConsFunc) Position() int {
	return op.BaseNode.Position
}

func (op *ConsFunc) Children() []Node {
	return op.entries
}

func (op *ConsFunc) SyntaxCheck() error {
	// Only 2 arguments for cons
	if len(op.entries) != 2 {
		return fmt.Errorf("cons operator requires 2 arguments, line %d, position %d", op.Line(), op.Position())
	}
	return nil
}

func (op *ConsFunc) Run(args []Node, env *Env) (Node, error) {
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
