package parser

import "fmt"

// A List

type LambdaOp struct {
	BaseNode
	entries []Node
}

func NewLambdaOp(line int, position int) *LambdaOp {
	return &LambdaOp{BaseNode: BaseNode{line, position}}
}

func (op *LambdaOp) String() string {
	return "Lambda Operator"
}

func (op *LambdaOp) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *LambdaOp) Line() int {
	return op.BaseNode.Line
}

func (op *LambdaOp) Position() int {
	return op.BaseNode.Position
}

func (op *LambdaOp) Children() []Node {
	return op.entries
}

func (op *LambdaOp) SyntaxCheck() error {
	if len(op.entries) != 2 {
		return fmt.Errorf("lambda op requires 2 arguments, line %d position %d", op.Line(), op.Position())
	}

	// first arg is a list of atoms
	argsNode := op.entries[0]

	if _, ok := argsNode.(*List); !ok {
		return fmt.Errorf("lambda op first arg must be a list, line %d position %d", op.Line(), op.Position())
	}

	allAtoms := true
	for _, n := range argsNode.Children() {
		if _, ok := n.(*Atom); !ok {
			allAtoms = false
		}
	}

	if !allAtoms {
		return fmt.Errorf("lambda op, arg list must all be atoms, line %d position %d", argsNode.Line(), argsNode.Position())
	}

	return nil
}

func (op *LambdaOp) Evaluate() (Node, error) {
	return nil, nil
}
