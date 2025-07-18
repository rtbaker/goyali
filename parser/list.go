package parser

// A List

type List struct {
	entries []Node
}

// Interface Node
func (list *List) Children() []Node {
	return list.entries
}

func (list *List) Evaluate() (Node, error) {
	return nil, nil
}
