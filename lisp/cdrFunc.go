package lisp

import "fmt"

// A List

type CdrFunc struct {
	BaseNode
	entries []Node
}

func NewCdrOp(line int, position int) *CdrFunc {
	return &CdrFunc{BaseNode: BaseNode{line, position}}
}

func (op *CdrFunc) String() string {
	return "Cdr Operator"
}

func (op *CdrFunc) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *CdrFunc) NodeType() string {
	return "Cdr Function"
}

func (op *CdrFunc) QuotedValue() Node {
	return NewAtom("cdr", op.Line(), op.Position())
}

func (op *CdrFunc) Line() int {
	return op.BaseNode.Line
}

func (op *CdrFunc) Position() int {
	return op.BaseNode.Position
}

func (op *CdrFunc) Children() []Node {
	return op.entries
}

func (op *CdrFunc) SyntaxCheck() error {
	// Only one argument for cdr
	if len(op.entries) != 1 {
		return fmt.Errorf("cdr operator requires only 1 argument, line %d, position %d", op.Line(), op.Position())
	}
	return nil
}

func (op *CdrFunc) Run(args []Node, env *Env) (Node, error) {
	// Only one argument for quote
	if len(args) != 1 {
		return nil, fmt.Errorf("cdr operator requires only 1 argument")
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

		retList := NewList(0, 0)

		for _, node := range listNode.Children()[1:] {
			retList.AppendNode(node)
		}

		return retList, nil
	}

	return nil, fmt.Errorf("cdr operator requires a list as its argument")
}
