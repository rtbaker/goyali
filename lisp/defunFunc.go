package lisp

import "fmt"

type DefunFunc struct {
	BaseNode
}

func NewDefunOp(line int, position int) *DefunFunc {
	return &DefunFunc{BaseNode: BaseNode{line, position}}
}

func (op *DefunFunc) String() string {
	return "Defun Operator"
}

// Interface Node
func (op *DefunFunc) NodeType() string {
	return "Defun Function"
}

func (op *DefunFunc) Line() int {
	return op.BaseNode.Line
}

func (op *DefunFunc) Position() int {
	return op.BaseNode.Position
}

func (op *DefunFunc) Run(args []Node, env *Env) (Node, error) {
	// Only one argument for quote
	if len(args) != 3 {
		return nil, fmt.Errorf("defun operator requires 2 arguments")
	}

	var labelAtom *Atom
	var ok bool

	if labelAtom, ok = args[0].(*Atom); !ok {
		return nil, fmt.Errorf("defun op expects first argument to be an atom")
	}

	userFunc, err := NewUserDefinedFunc(args[1], args[2])

	if err != nil {
		return nil, err
	}

	env.addSymbol(labelAtom.Name, userFunc)

	return NilAtom(), nil
}
