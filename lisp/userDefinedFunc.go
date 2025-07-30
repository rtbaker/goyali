package lisp

// A List

type UserDefinedFunc struct {
	BaseNode
	args       []Node
	expression Node
}

func NewUserDefinedFunc(args []Node, expression Node) *UserDefinedFunc {
	return &UserDefinedFunc{
		BaseNode:   BaseNode{0, 0},
		args:       args,
		expression: expression,
	}
}

func (op *UserDefinedFunc) String() string {
	return "Function"
}

// Interface Node
func (op *UserDefinedFunc) NodeType() string {
	return "User Defined Function"
}

func (op *UserDefinedFunc) Line() int {
	return op.BaseNode.Line
}

func (op *UserDefinedFunc) Position() int {
	return op.BaseNode.Position
}

func (op *UserDefinedFunc) Run(args []Node, env *Env) (Node, error) {
	return nil, nil
}
