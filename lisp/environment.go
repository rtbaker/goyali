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
		return env.prev.getSymbol(name)
	}

	return node
}
