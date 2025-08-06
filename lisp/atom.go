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
	return atom.Name
}

// Interface Node
func (atom *Atom) NodeType() string {
	return "Atom"
}

func (atom *Atom) Line() int {
	return atom.BaseNode.Line
}

func (atom *Atom) Position() int {
	return atom.BaseNode.Position
}

func (atom *Atom) Evaluate(env *Env, inQuote bool) (Node, error) {
	if inQuote {
		return atom, nil
	}

	retNode := env.getSymbol(atom.Name)

	if retNode == nil {
		return nil, NewLispError(fmt.Sprintf("Atom \"%s\" has no value", atom.Name), atom.Line(), atom.Position(), nil)
	}

	return retNode, nil
}
