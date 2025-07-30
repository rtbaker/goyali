package lisp

import "fmt"

// A List

type LambdaFunc struct {
	BaseNode
}

func NewLambdaOp(line int, position int) *LambdaFunc {
	return &LambdaFunc{BaseNode: BaseNode{line, position}}
}

func (op *LambdaFunc) String() string {
	return "Lambda Operator"
}

// Interface Node
func (op *LambdaFunc) NodeType() string {
	return "Lambda Function"
}

func (op *LambdaFunc) Line() int {
	return op.BaseNode.Line
}

func (op *LambdaFunc) Position() int {
	return op.BaseNode.Position
}

func (op *LambdaFunc) Run(args []Node, env *Env) (Node, error) {
	// Only one argument for quote
	if len(args) != 2 {
		return nil, fmt.Errorf("lambda operator requires 2 arguments")
	}

	return NewUserDefinedFunc(args[0], args[1])
}
