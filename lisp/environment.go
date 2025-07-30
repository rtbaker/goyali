package lisp

// Runtime environment (symbol tables)

type Env struct {
	symbols map[string]Node
	prev    *Env
}

func NewEnv(prev *Env) *Env {
	return &Env{
		symbols: make(map[string]Node),
		prev:    prev,
	}
}
func (env *Env) addSymbol(name string, n Node) {
	env.symbols[name] = n
}

func (env *Env) getSymbol(name string) Node {
	var node Node
	var ok bool

	if node, ok = env.symbols[name]; !ok {
		if env.prev == nil {
			return nil
		}

		return env.prev.getSymbol(name)
	}

	return node
}

func (env *Env) initialiseBuiltin() {
	// Add the builtin functions here
	env.addSymbol("quote", NewQuoteOp(0, 0))
	env.addSymbol("atom", NewAtomOp(0, 0))
	env.addSymbol("eq", NewEqualsOp(0, 0))
	env.addSymbol("car", NewCarOp(0, 0))
	env.addSymbol("cdr", NewCdrOp(0, 0))
	env.addSymbol("cons", NewConsOp(0, 0))
}
