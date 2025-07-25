package lisp

import "fmt"

type DefunOp struct {
	BaseNode
	entries []Node
}

func NewDefunOp(line int, position int) *DefunOp {
	return &DefunOp{BaseNode: BaseNode{line, position}}
}

func (op *DefunOp) String() string {
	return "Defun Operator"
}

func (op *DefunOp) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *DefunOp) QuotedValue() Node {
	return NewAtom("defun", op.Line(), op.Position())
}

func (op *DefunOp) Line() int {
	return op.BaseNode.Line
}

func (op *DefunOp) Position() int {
	return op.BaseNode.Position
}

func (op *DefunOp) Children() []Node {
	return op.entries
}

func (op *DefunOp) SyntaxCheck() error {
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

	allAtoms := true
	for _, n := range argsNode.Children() {
		if _, ok := n.(*Atom); !ok {
			allAtoms = false
		}
	}

	if !allAtoms {
		return fmt.Errorf("defun op, arg list must all be atoms, line %d position %d", argsNode.Line(), argsNode.Position())
	}

	return nil
}

func (op *DefunOp) Evaluate() (Node, error) {
	return nil, nil
}
