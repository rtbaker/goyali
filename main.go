package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/chzyer/readline"
	"github.com/rtbaker/goyali/lexer"
	"github.com/rtbaker/goyali/lisp"
)

var usage string = `%s [OPTION] [FILE]
Run lisp interpreter on FILE or if FILE not specified run interactively.

Options
  --help			Show this message.
  --lib=DIRECTORY		Load all *.lisp files in DIRECTORY before running.
`

func main() {
	var libraryDir string
	flag.StringVar(&libraryDir, "lib", "", "Directory to preload *.lisp files from")

	var showUsage bool
	flag.BoolVar(&showUsage, "help", false, "Show the usage message")

	flag.Parse()

	if showUsage {
		fmt.Printf(usage, os.Args[0])
		os.Exit(0)
	}

	// Setup top level env/symbol table
	env := lisp.NewEnv(nil)
	env.InitialiseBuiltin()

	if libraryDir != "" {
		err := preloadDirectory(env, libraryDir)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	remainingArgs := len(flag.Args())

	if remainingArgs > 1 {
		fmt.Printf(usage, os.Args[0])
		os.Exit(1)
	}

	if remainingArgs == 1 {
		err := runFile(env, flag.Args()[0])

		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
	} else {
		runInteractive(env)
	}

	os.Exit(0)
}

func runFile(env *lisp.Env, filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("no such file: %s", filename)
	}

	dat, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	lex := lexer.NewLexer(strings.NewReader(string(dat)))

	myParser := lisp.NewParser(lex)
	program, err := myParser.ParseProgram()

	if err != nil {
		return err
	}

	for _, expression := range program.Expressions {
		node, err := lisp.EvaluateNode(expression, env, false)

		if err != nil {
			return err
		}

		fmt.Printf("%s\n", node)
	}

	return nil
}

func preloadDirectory(env *lisp.Env, directory string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return fmt.Errorf("library directory does not exist")
	}

	path := filepath.Join(directory, "*.lisp")
	matches, err := filepath.Glob(path)

	if err != nil {
		return err
	}

	for _, file := range matches {
		err = preloadFile(env, file)

		if err != nil {
			return fmt.Errorf("Error preloading library file %s: %s", file, err)
		}
	}

	return nil
}

func preloadFile(env *lisp.Env, filename string) error {
	dat, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	lex := lexer.NewLexer(strings.NewReader(string(dat)))

	myParser := lisp.NewParser(lex)
	program, err := myParser.ParseProgram()

	if err != nil {
		return err
	}

	for _, expression := range program.Expressions {
		_, err := lisp.EvaluateNode(expression, env, false)

		if err != nil {
			return err
		}
	}

	return nil
}

func runInteractive(env *lisp.Env) {
	var node lisp.Node
	var err error

	//node, err = myParser.GetExpression()

	var builder strings.Builder

	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		text, err := rl.Readline()
		if err != nil { // io.EOF
			break
		}

		builder.WriteString(text)

		lex := lexer.NewLexer(strings.NewReader(builder.String()))

		myParser := lisp.NewParser(lex)
		node, err = myParser.GetExpression()

		if err != nil {
			fmt.Printf("error: %s\n", err)
			builder.Reset()
			lex.ResetLineNos()
			continue
		}

		var resultNode lisp.Node

		if node != nil {
			resultNode, err = lisp.EvaluateNode(node, env, false)

			if err != nil {
				fmt.Printf("error: %s\n", err)
			} else {
				fmt.Printf("%s\n", resultNode)
			}

			builder.Reset()
			lex.ResetLineNos()
		}
	}
}
