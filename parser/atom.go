package parser

import "fmt"

// A basic atom
type Atom struct {
	Value string
}

func NewAtom(val string) *Atom {
	return &Atom{
		Value: val,
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
