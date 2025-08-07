package lisp

// A List

type QuoteOp struct {
	BaseNode
}

func NewQuoteOp(line int, position int) *QuoteOp {
	return &QuoteOp{BaseNode: BaseNode{line, position}}
}

func (op *QuoteOp) String() string {
	return "Quote Operator"
}

// Interface Node
func (op *QuoteOp) NodeType() string {
	return "Quote Function"
}

func (op *QuoteOp) Line() int {
	return op.BaseNode.Line
}

func (op *QuoteOp) Position() int {
	return op.BaseNode.Position
}

func (op *QuoteOp) Run(args []Node, env *Env) (Node, error) {
	// Only one argument for quote
	if len(args) != 1 {
		return nil, NewSimpleLispError("quote operator requires only 1 argument")
	}

	retNode, err := EvaluateNode(args[0], env, true)

	if err != nil {
		return nil, err
	}

	return retNode, nil
}
