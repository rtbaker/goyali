package lisp

import "fmt"

// A List

type LambdaFunc struct {
	BaseNode
	entries []Node
}

func NewLambdaOp(line int, position int) *LambdaFunc {
	return &LambdaFunc{BaseNode: BaseNode{line, position}}
}

func (op *LambdaFunc) String() string {
	return "Lambda Operator"
}

func (op *LambdaFunc) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *LambdaFunc) NodeType() string {
	return "Lambda Function"
}

func (op *LambdaFunc) QuotedValue() Node {
	return NewAtom("lambda", op.Line(), op.Position())
}

func (op *LambdaFunc) Line() int {
	return op.BaseNode.Line
}

func (op *LambdaFunc) Position() int {
	return op.BaseNode.Position
}

func (op *LambdaFunc) Children() []Node {
	return op.entries
}

func (op *LambdaFunc) SyntaxCheck() error {
	if len(op.entries) != 2 {
		return fmt.Errorf("lambda op requires 2 arguments, line %d position %d", op.Line(), op.Position())
	}

	// first arg is a list of atoms
	argsNode := op.entries[0]

	if _, ok := argsNode.(*List); !ok {
		return fmt.Errorf("lambda op first arg must be a list, line %d position %d", op.Line(), op.Position())
	}

	/*
		allAtoms := true
		for _, n := range argsNode.Children() {
			if _, ok := n.(*Atom); !ok {
				allAtoms = false
			}
		}

		if !allAtoms {
			return fmt.Errorf("lambda op, arg list must all be atoms, line %d position %d", argsNode.Line(), argsNode.Position())
		}
	*/
	return nil
}

func (op *LambdaFunc) Run(args []Node) (Node, error) {
	return nil, nil
}
