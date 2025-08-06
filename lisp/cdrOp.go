package lisp

// A List

type CdrOp struct {
	BaseNode
}

func NewCdrOp(line int, position int) *CdrOp {
	return &CdrOp{BaseNode: BaseNode{line, position}}
}

func (op *CdrOp) String() string {
	return "Cdr Operator"
}

// Interface Node
func (op *CdrOp) NodeType() string {
	return "Cdr Function"
}

func (op *CdrOp) Line() int {
	return op.BaseNode.Line
}

func (op *CdrOp) Position() int {
	return op.BaseNode.Position
}

func (op *CdrOp) Run(args []Node, env *Env) (Node, error) {
	// Only one argument for quote
	if len(args) != 1 {
		return nil, NewSimpleLispError("cdr operator requires only 1 argument")
	}

	retNode, err := EvaluateNode(args[0], env, false)

	if err != nil {
		return nil, err
	}

	if listNode, ok := retNode.(*List); ok {
		// Empty list, returns as nil
		if len(listNode.Children()) <= 1 {
			return NewList(0, 0), nil
		}

		retList := NewList(0, 0)

		for _, node := range listNode.Children()[1:] {
			retList.AppendNode(node)
		}

		return retList, nil
	}

	return nil, NewSimpleLispError("cdr operator requires a list as its argument")
}
