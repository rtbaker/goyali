package lisp

// Interface for a basic node in the AST
type Node interface {
	// Get a human readable explaintion of the node type
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
}

// Common parts of a Concrete Node type
type BaseNode struct {
	Line     int
	Position int
}
