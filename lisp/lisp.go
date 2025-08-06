package lisp

import "fmt"

func RunProgram(prog *Program) error {
	for _, expression := range prog.Expressions {
		retNode, err := EvaluateNode(expression, prog.env, false)

		if err != nil {
			return err
		}

		if !IsNil(retNode) {
			fmt.Printf("%s\n", retNode)
		}
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

	return nil, NewLispError(fmt.Sprintf("cannot evaluate a %s node", node.NodeType()), node.Line(), node.Position(), nil)
}
