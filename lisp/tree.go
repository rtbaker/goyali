package lisp

// Stuff for walking the tree

func WalkTreeSimple(top Node, f func(n Node) error) error {
	err := f(top)
	if err != nil {
		return err
	}

	for _, n := range top.Children() {
		err := WalkTreeSimple(n, f)
		if err != nil {
			return err
		}
	}

	return nil
}

func WalkTree(top Node, f func(n Node) error, preChildren func() error, postChildren func() error) error {
	err := f(top)
	if err != nil {
		return err
	}

	preChildren()

	for _, n := range top.Children() {
		err := WalkTree(n, f, preChildren, postChildren)
		if err != nil {
			return err
		}
	}

	postChildren()

	return nil
}

func SyntaxCheckTree(top Node) error {
	return WalkTreeSimple(top, func(n Node) error {
		return n.SyntaxCheck()
	})
}
