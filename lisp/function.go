package lisp

// A function is something that can do work, will be the first
// thing in a list that is being eval'ed and not quoted

type function interface {
	Run(args []Node) (Node, error)
}
