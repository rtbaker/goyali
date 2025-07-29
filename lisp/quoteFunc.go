package lisp

import "fmt"

// A List

type QuoteFunc struct {
	BaseNode
	entries []Node
}

func NewQuoteOp(line int, position int) *QuoteFunc {
	return &QuoteFunc{BaseNode: BaseNode{line, position}}
}

func (op *QuoteFunc) String() string {
	return "Quote Operator"
}

func (op *QuoteFunc) AppendNode(n Node) {
	op.entries = append(op.entries, n)
}

// Interface Node
func (op *QuoteFunc) NodeType() string {
	return "Quote Function"
}

func (op *QuoteFunc) Line() int {
	return op.BaseNode.Line
}

func (op *QuoteFunc) Position() int {
	return op.BaseNode.Position
}

func (op *QuoteFunc) Run(args []Node, env *Env) (Node, error) {
	// Only one argument for quote
	if len(args) != 1 {
		return nil, fmt.Errorf("quote operator requires only 1 argument")
	}

	retNode, err := EvaluateNode(args[0], env, true)

	if err != nil {
		return nil, err
	}

	return retNode, nil
}
