package lisp

// A List

type AtomOp struct {
	BaseNode
}

func NewAtomOp(line int, position int) *AtomOp {
	return &AtomOp{BaseNode: BaseNode{line, position}}
}

func (op *AtomOp) String() string {
	return "Atom Operator"
}

// Interface Node
func (op *AtomOp) NodeType() string {
	return "Atom Function"
}

func (op *AtomOp) Line() int {
	return op.BaseNode.Line
}

func (op *AtomOp) Position() int {
	return op.BaseNode.Position
}

func (op *AtomOp) Run(args []Node, env *Env) (Node, error) {
	// Only one argument for quote
	if len(args) != 1 {
		return nil, NewLispError("atom operator requires only 1 argument", op.Line(), op.Position(), nil)
	}

	retNode, err := EvaluateNode(args[0], env, false)

	if err != nil {
		return nil, err
	}

	if NodeIsAtom(retNode) {
		return Truth(), nil
	}

	if listVal, ok := retNode.(*List); ok {
		if listVal.isEmptyList() {
			return Truth(), nil
		}
	}

	return Falsity(), nil
}
