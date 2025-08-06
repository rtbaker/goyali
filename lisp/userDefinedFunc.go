package lisp

import "fmt"

// A List

type UserDefinedFunc struct {
	BaseNode
	args       []Node
	expression Node
}

func NewUserDefinedFunc(args Node, expression Node) (*UserDefinedFunc, error) {
	userFunc := &UserDefinedFunc{
		BaseNode:   BaseNode{0, 0},
		expression: expression,
	}

	// Args must be a list atoms
	if argList, ok := args.(*List); ok {
		for i, child := range argList.Children() {
			if _, ok := child.(*Atom); !ok {
				return nil, fmt.Errorf("function arguments must be atoms (argument %d is not)", i)
			}
		}

		userFunc.args = argList.Children()
	} else {
		return nil, fmt.Errorf("function args must be a list")
	}

	return userFunc, nil
}

func (op *UserDefinedFunc) String() string {
	return "Function"
}

// Interface Node
func (op *UserDefinedFunc) NodeType() string {
	return "User Defined Function"
}

func (op *UserDefinedFunc) Line() int {
	return op.BaseNode.Line
}

func (op *UserDefinedFunc) Position() int {
	return op.BaseNode.Position
}

func (op *UserDefinedFunc) Run(args []Node, env *Env) (Node, error) {
	if len(args) != len(op.args) {
		return nil, fmt.Errorf("wring number of arguments to function, %d required, %d given", len(op.args), len(args))
	}

	// create the symbol table for this run
	funcEnv := NewEnv(env)

	for i, requiredArg := range op.args {
		argValue, err := EvaluateNode(args[i], funcEnv, false)

		if err != nil {
			return nil, fmt.Errorf("error evaluating argument to function (argument %d): %s", i, err)
		}

		// we check these were atoms so no need to check again
		argAtom := requiredArg.(*Atom)
		funcEnv.addSymbol(argAtom.Name, argValue)
	}

	node, err := EvaluateNode(op.expression, funcEnv, false)

	return node, err
}
