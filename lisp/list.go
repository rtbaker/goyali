package lisp

// A List

type List struct {
	BaseNode
	entries []Node
}

func NewList(line int, position int) *List {
	return &List{BaseNode: BaseNode{line, position}}
}

func (list *List) String() string {
	if len(list.entries) > 0 {
		return "List"
	}

	return "Empty List"
}

func (list *List) isEmptyList() bool {
	return len(list.entries) == 0
}

func (list *List) AppendNode(n Node) {
	list.entries = append(list.entries, n)
}

// Interface Node
func (list *List) Line() int {
	return list.BaseNode.Line
}

func (list *List) Position() int {
	return list.BaseNode.Position
}

func (list *List) Children() []Node {
	return list.entries
}

func (list *List) SyntaxCheck() error {
	return nil
}

func (list *List) Evaluate() (Node, error) {
	return nil, nil
}
