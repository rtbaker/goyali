package lisp

import "fmt"

// A basic atom
type Atom struct {
	BaseNode
	Name string
}

func NewAtom(name string, line int, position int) *Atom {
	return &Atom{
		BaseNode: BaseNode{line, position},
		Name:     name,
	}
}

func (atom *Atom) String() string {
	return fmt.Sprintf("Atom: %s", atom.Name)
}

// Interface Node
func (atom *Atom) QuotedValue() Node {
	return atom
}

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
	if len(atom.Name) == 0 {
		// Not sure how this would happen but it's not right
		return fmt.Errorf("zero length atom value, line %d position %d", atom.Line(), atom.Position())
	}

	return nil
}

func (atom *Atom) Evaluate() (Node, error) {
	return nil, nil
}
