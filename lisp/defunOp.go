package lisp

type DefunOp struct {
	BaseNode
}

func NewDefunOp(line int, position int) *DefunOp {
	return &DefunOp{BaseNode: BaseNode{line, position}}
}

func (op *DefunOp) String() string {
	return "Defun Operator"
}

// Interface Node
func (op *DefunOp) NodeType() string {
	return "Defun Function"
}

func (op *DefunOp) Line() int {
	return op.BaseNode.Line
}

func (op *DefunOp) Position() int {
	return op.BaseNode.Position
}

func (op *DefunOp) Run(args []Node, env *Env) (Node, error) {
	if len(args) != 3 {
		return nil, NewSimpleLispError("defun operator requires 3 arguments")
	}

	var labelAtom *Atom
	var ok bool

	if labelAtom, ok = args[0].(*Atom); !ok {
		return nil, NewSimpleLispError("defun op expects first argument to be an atom")
	}

	userFunc, err := NewUserDefinedFunc(args[1], args[2])

	if err != nil {
		return nil, err
	}

	env.addSymbol(labelAtom.Name, userFunc)

	return NilAtom(), nil
}
