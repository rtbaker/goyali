package parser

// A single expression

type Expression struct {
}

// Interface Node
func (expr *Expression) Children() []Node {
	return nil // No implemented
}

func (expr *Expression) Evaluate() (Node, error) {
	return nil, nil
}
