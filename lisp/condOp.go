package lisp

import "fmt"

// A List

type CondOp struct {
	BaseNode
}

func NewCondOp(line int, position int) *CondOp {
	return &CondOp{BaseNode: BaseNode{line, position}}
}

func (op *CondOp) String() string {
	return "Cond Operator"
}

// Interface Node
func (op *CondOp) NodeType() string {
	return "Cond Function"
}

func (op *CondOp) Line() int {
	return op.BaseNode.Line
}

func (op *CondOp) Position() int {
	return op.BaseNode.Position
}

func (op *CondOp) Run(args []Node, env *Env) (Node, error) {
	for _, arg := range args {
		// each argument is a list of 2 things, a test and an expression
		var listArg *List
		var ok bool

		if listArg, ok = arg.(*List); !ok {
			return nil, fmt.Errorf("argument to cond must be a list, line %d, position %d", arg.Line(), arg.Position())
		}

		if len(listArg.Children()) != 2 {
			return nil, fmt.Errorf("cond argument must be a list of 2 items, test and expression, line %d, position %d", arg.Line(), arg.Position())
		}

		retNode1, err := EvaluateNode(listArg.Children()[0], env, false)
		if err != nil {
			return nil, err
		}

		if !IsTrue(retNode1) {
			continue
		}

		retNode2, err := EvaluateNode(listArg.Children()[1], env, false)

		if err != nil {
			return nil, err
		}

		return retNode2, nil
	}

	// no match, return a "NIL" atom
	return NilAtom(), nil
}
