package lisp

import "fmt"

// A List

type LambdaOp struct {
	BaseNode
}

func NewLambdaOp(line int, position int) *LambdaOp {
	return &LambdaOp{BaseNode: BaseNode{line, position}}
}

func (op *LambdaOp) String() string {
	return "Lambda Operator"
}

// Interface Node
func (op *LambdaOp) NodeType() string {
	return "Lambda Function"
}

func (op *LambdaOp) Line() int {
	return op.BaseNode.Line
}

func (op *LambdaOp) Position() int {
	return op.BaseNode.Position
}

func (op *LambdaOp) Run(args []Node, env *Env) (Node, error) {
	// Only one argument for quote
	if len(args) != 2 {
		return nil, fmt.Errorf("lambda operator requires 2 arguments")
	}

	return NewUserDefinedFunc(args[0], args[1])
}
