package lisp

import "fmt"

// A List

type LabelOp struct {
	BaseNode
}

func NewLabelOp(line int, position int) *LabelOp {
	return &LabelOp{BaseNode: BaseNode{line, position}}
}

func (op *LabelOp) String() string {
	return "Label Operator"
}

// Interface Node
func (op *LabelOp) NodeType() string {
	return "Label Function"
}

func (op *LabelOp) Line() int {
	return op.BaseNode.Line
}

func (op *LabelOp) Position() int {
	return op.BaseNode.Position
}

func (op *LabelOp) Run(args []Node, env *Env) (Node, error) {
	// Only one argument for quote
	if len(args) != 2 {
		return nil, fmt.Errorf("label operator requires 2 arguments")
	}

	var labelAtom *Atom
	var ok bool

	if labelAtom, ok = args[0].(*Atom); !ok {
		return nil, fmt.Errorf("label op expects first argument to be an atom")
	}

	expr, err := EvaluateNode(args[1], env, false)

	if err != nil {
		return nil, fmt.Errorf("error evaluating 2nd argument to label: %s", err)
	}

	if _, ok := expr.(LispFunction); !ok {
		return nil, fmt.Errorf("label op expects 2nd argument to be a function")
	}

	env.addSymbol(labelAtom.Name, expr)
	return NilAtom(), nil
}
