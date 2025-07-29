package lisp

import "fmt"

func runProgram(prog *Program) error {
	for _, expression := range prog.Expressions {
		retNode, err := EvaluateNode(expression, false)

		if err != nil {
			return err
		}

		fmt.Printf("%s\n", retNode)
	}

	return nil
}

func EvaluateNode(node Node, inQuote bool) (Node, error) {
	//if listNode, ok := node.(ListNode); ok {

	//}

	return nil, nil
}
