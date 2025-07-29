package lisp

import "fmt"

func RunProgram(prog *Program) error {
	for _, expression := range prog.Expressions {
		retNode, err := EvaluateNode(expression, prog.env, false)

		if err != nil {
			return err
		}

		fmt.Printf("%s\n", retNode)
	}

	return nil
}

func EvaluateNode(node Node, env *Env, inQuote bool) (Node, error) {
	if node, ok := node.(EvaluatableNode); ok {
		resultNode, err := node.Evaluate(env, inQuote)

		if err != nil {
			return nil, err
		}

		return resultNode, nil
	}

	return nil, fmt.Errorf("cannot evaluate a %s node at line %d, position %d", node.NodeType(), node.Line(), node.Position())
}
