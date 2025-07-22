package parser

import "fmt"

// A basic atom
type Atom struct {
	BaseNode
	Value string
}

func NewAtom(val string, line int, position int) *Atom {
	return &Atom{
		BaseNode: BaseNode{line, position},
		Value:    val,
	}
}

func (atom *Atom) String() string {
	return fmt.Sprintf("Atom: %s", atom.Value)
}

// Interface Node
func (atom *Atom) Children() []Node {
	return nil // No children
}

func (atom *Atom) Evaluate() (Node, error) {
	return nil, nil
}
