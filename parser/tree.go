package parser

// Stuff for walking the tree

func WalkTree(top Node, f func(n Node) error) error {
	err := f(top)
	if err != nil {
		return err
	}

	for _, n := range top.Children() {
		err := WalkTree(n, f)
		if err != nil {
			return err
		}
	}

	return nil
}

func SyntaxCheckTree(top Node) error {
	return WalkTree(top, func(n Node) error {
		return n.SyntaxCheck()
	})
}
