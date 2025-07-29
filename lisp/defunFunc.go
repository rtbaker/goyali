package lisp

import "fmt"

type DefunFunc struct {
	BaseNode
	entries []Node
}

func NewDefunOp(line int, position int) *DefunFunc {
	return &DefunFunc{BaseNode: BaseNode{line, position}}
}

func (op *DefunFunc) String() string {
	return "Defun Operator"
}

func (op *DefunFunc) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *DefunFunc) NodeType() string {
	return "Defun Function"
}

func (op *DefunFunc) QuotedValue() Node {
	return NewAtom("defun", op.Line(), op.Position())
}

func (op *DefunFunc) Line() int {
	return op.BaseNode.Line
}

func (op *DefunFunc) Position() int {
	return op.BaseNode.Position
}

func (op *DefunFunc) Children() []Node {
	return op.entries
}

func (op *DefunFunc) SyntaxCheck() error {
	if len(op.entries) != 3 {
		return fmt.Errorf("defun op requires 3 arguments, line %d position %d", op.Line(), op.Position())
	}

	// First arg is an atom
	labelNode := op.entries[0]

	if _, ok := labelNode.(*Atom); !ok {
		return fmt.Errorf("defun op first arg must be an atom, line %d position %d", op.Line(), op.Position())
	}

	// second arg is a list of atoms
	argsNode := op.entries[1]

	if _, ok := argsNode.(*List); !ok {
		return fmt.Errorf("defun op second arg must be a list, line %d position %d", op.Line(), op.Position())
	}
	/*
		allAtoms := true
		for _, n := range argsNode.Children() {
			if _, ok := n.(*Atom); !ok {
				allAtoms = false
			}
		}

		if !allAtoms {
			return fmt.Errorf("defun op, arg list must all be atoms, line %d position %d", argsNode.Line(), argsNode.Position())
		}
	*/
	return nil
}

func (op *DefunFunc) Run(args []Node) (Node, error) {
	return nil, nil
}
