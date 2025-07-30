package lisp

// Interface for a basic node in the AST
type Node interface {
	// Get a human readable explanation of the node type
	NodeType() string
	// What line of the source was this on?
	Line() int
	// Position
	Position() int
}

// List type node
type ListNode interface {
	Node
	// Used to walk the tree
	Children() []Node
	AppendNode(n Node)
	AppendNodes(nodes []Node)
}

// Can this node be evaluated?
type EvaluatableNode interface {
	Node
	Evaluate(env *Env, inQuote bool) (Node, error)
}

// Common parts of a Concrete Node type
type BaseNode struct {
	Line     int
	Position int
}

func NodeIsAtom(n Node) bool {
	_, ok := n.(*Atom)
	return ok
}

func Truth() Node {
	return NewAtom("t", 0, 0)
}

func Falsity() Node {
	return NewList(0, 0) // empty list is false
}
