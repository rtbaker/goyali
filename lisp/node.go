package lisp

// Interface for a basic node in the AST
type Node interface {
	// Evaluate (implicitly walks children), returns a pointer to a Node type as the result (or nil)
	Evaluate() (Node, error)
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
}

// Common parts of a Concrete Node type
type BaseNode struct {
	Line     int
	Position int
}
