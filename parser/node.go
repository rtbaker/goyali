package parser

// Interface for a node in the AST

type Node interface {
	// Used to walk the tree
	Children() []Node
	// Evaluate (implicitly walks children), returns a pointer to a Node type as the result (or nil)
	Evaluate() (Node, error)
}
