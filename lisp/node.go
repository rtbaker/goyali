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

// Misc' routines for dealing with basic node things
func NodeIsAtom(n Node) bool {
	_, ok := n.(*Atom)
	return ok
}

func Truth() Node {
	return NewAtom("t", 0, 0)
}

func IsTrue(n Node) bool {
	if nodeAtom, ok := n.(*Atom); ok {
		return nodeAtom.Name == "t"
	}

	return false
}

func Falsity() Node {
	return NewList(0, 0) // empty list is false
}

func IsFalse(n Node) bool {
	return IsEmptyList(n)
}

func IsEmptyList(n Node) bool {
	if nodeList, ok := n.(*List); ok {
		return len(nodeList.entries) == 0
	}

	return false
}

func NilAtom() Node {
	return NewAtom("NIL", 0, 0)
}

func IsNil(n Node) bool {
	if nodeAtom, ok := n.(*Atom); ok {
		return nodeAtom.Name == "NIL"
	}

	return false
}
