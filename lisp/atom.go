package lisp

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
func (atom *Atom) Line() int {
	return atom.BaseNode.Line
}

func (atom *Atom) Position() int {
	return atom.BaseNode.Position
}

func (atom *Atom) Children() []Node {
	return nil // No children
}

func (atom *Atom) SyntaxCheck() error {
	if len(atom.Value) == 0 {
		// Not sure how this would happen but it's not right
		return fmt.Errorf("zero length atom value, line %d position %d", atom.Line(), atom.Position())
	}

	return nil
}

func (atom *Atom) Evaluate() (Node, error) {
	return nil, nil
}
