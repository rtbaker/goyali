package lisp

// Stuff for walking the tree

func WalkTreeSimple(top ListNode, f func(n Node) error) error {
	err := f(top)
	if err != nil {
		return err
	}

	for _, n := range top.Children() {
		if listNode, ok := n.(ListNode); ok {
			err := WalkTreeSimple(listNode, f)
			if err != nil {
				return err
			}
		} else {
			f(n)
		}
	}

	return nil
}

func WalkTree(top ListNode, f func(n Node) error, preChildren func() error, postChildren func() error) error {
	err := f(top)
	if err != nil {
		return err
	}

	preChildren()

	for _, n := range top.Children() {
		if listNode, ok := n.(ListNode); ok {
			err := WalkTree(listNode, f, preChildren, postChildren)
			if err != nil {
				return err
			}
		} else {
			f(n)
		}
	}

	postChildren()

	return nil
}

/*
func SyntaxCheckTree(top Node) error {
	return WalkTreeSimple(top, func(n Node) error {
		return n.SyntaxCheck()
	})
}
*/
