package parser

// Interface for a basic node in the AST
type Node interface {
	// Used to walk the tree
	Children() []Node
	// Is the node well formed?
	SyntaxCheck() error
	// Evaluate (implicitly walks children), returns a pointer to a Node type as the result (or nil)
	Evaluate() (Node, error)
}

// List type node
type ListNode interface {
	Node
	AppendNode(n Node)
}

// Common parts of a Concrete Node type
type BaseNode struct {
	Line     int
	Position int
}
