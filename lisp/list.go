package lisp

import (
	"fmt"
	"strings"
)

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
		var b strings.Builder
		b.WriteString("(")

		for i, child := range list.entries {
			if i > 0 {
				b.WriteString(" ")
			}
			b.WriteString(fmt.Sprintf("%s", child))
		}

		b.WriteString(")")
		return b.String()
	}

	return "()"
}

func (list *List) isEmptyList() bool {
	return len(list.entries) == 0
}

func (list *List) AppendNode(n Node) {
	list.entries = append(list.entries, n)
}

func (list *List) AppendNodes(nodes []Node) {
	list.entries = append(list.entries, nodes...)
}

// Interface Node
func (list *List) NodeType() string {
	return "List"
}

func (list *List) Line() int {
	return list.BaseNode.Line
}

func (list *List) Position() int {
	return list.BaseNode.Position
}

func (list *List) Children() []Node {
	return list.entries
}

func (list *List) Evaluate(env *Env, inQuote bool) (Node, error) {
	// empty/quoted list evaluates as itself
	if len(list.entries) == 0 || inQuote {
		return list, nil
	}

	// Otherwise the first child must be a function
	firstNode := list.entries[0]
	eFirstNode, err := EvaluateNode(firstNode, env, inQuote)

	if err != nil {
		return nil, NewLispError("evaluating first list item", firstNode.Line(), firstNode.Position(), err)
	}

	if lFunc, ok := eFirstNode.(LispFunction); ok {
		retNode, err := lFunc.Run(list.entries[1:], env)

		if err != nil {
			return nil, NewLispError(fmt.Sprintf("%s", eFirstNode.NodeType()), firstNode.Line(), firstNode.Position(), err)
		}

		return retNode, nil
	} else {
		return nil, NewLispError(fmt.Sprintf("undefined function %s", eFirstNode), firstNode.Line(), firstNode.Position(), nil)
	}
}
