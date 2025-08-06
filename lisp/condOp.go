package lisp

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
			return nil, NewLispError("argument to cond must be a list", arg.Line(), arg.Position(), nil)
		}

		if len(listArg.Children()) != 2 {
			return nil, NewLispError("cond argument must be a list of 2 items, test and expression", arg.Line(), arg.Position(), nil)
		}

		retNode1, err := EvaluateNode(listArg.Children()[0], env, false)
		if err != nil {
			return nil, NewLispError("cannot evaluate first argument", listArg.Children()[0].Line(), listArg.Children()[0].Position(), err)
		}

		if !IsTrue(retNode1) {
			continue
		}

		retNode2, err := EvaluateNode(listArg.Children()[1], env, false)

		if err != nil {
			return nil, NewLispError("cannot evaluate second argument", listArg.Children()[1].Line(), listArg.Children()[1].Position(), err)
		}

		return retNode2, nil
	}

	// no match, return a "NIL" atom
	return NilAtom(), nil
}
