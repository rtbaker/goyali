package lisp

import "fmt"

// A List

type CondFunc struct {
	BaseNode
	entries []Node
}

func NewCondOp(line int, position int) *CondFunc {
	return &CondFunc{BaseNode: BaseNode{line, position}}
}

func (op *CondFunc) String() string {
	return "Cond Operator"
}

func (op *CondFunc) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *CondFunc) NodeType() string {
	return "Cond Function"
}

func (op *CondFunc) QuotedValue() Node {
	return NewAtom("cond", op.Line(), op.Position())
}

func (op *CondFunc) Line() int {
	return op.BaseNode.Line
}

func (op *CondFunc) Position() int {
	return op.BaseNode.Position
}

func (op *CondFunc) Children() []Node {
	return op.entries
}

func (op *CondFunc) SyntaxCheck() error {
	// Must have at least 1 entry?
	if len(op.entries) == 0 {
		return fmt.Errorf("cond operator requires at least 1 argument, line %d position %d", op.Line(), op.Position())
	}

	// each entry is (<test expression> <return expression>)
	//	for _, n := range op.Children() {
	//		if len(n.Children()) != 2 {
	//	return fmt.Errorf("cond entry requires 2 expressions, line %d position %d", n.Line(), n.Position())
	//	}
	//	}
	return nil
}

func (op *CondFunc) Run(args []Node, env *Env) (Node, error) {
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
