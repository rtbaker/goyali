package lisp

// Top level, a list of expressions

type Program struct {
	env         *Env
	Expressions []Node
}

func NewProgram() *Program {
	// top level symbol table
	env := NewEnv(nil)
	env.InitialiseBuiltin()

	return &Program{
		env: env,
	}
}

func (prog *Program) String() string {
	return "Program"
}

// Interface Node
func (prog *Program) Line() int {
	return 0
}

func (prog *Program) Position() int {
	return 0
}

func (prog *Program) Children() []Node {
	return prog.Expressions
}

func (prog *Program) AppendNode(n Node) {
	prog.Expressions = append(prog.Expressions, n)
}
