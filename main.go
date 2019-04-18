package main

import (
	"bufio"
	"fmt"
	"strings"

	"bitbucket.org/rtbaker/goyali/lexer"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(strings.NewReader("23 + 34 hello"))
	lex := lexer.NewLexer(reader)

	fmt.Println("Start typing: ")

	var token *lexer.Token
	var err error

	for token, err = lex.GetToken(); token.Code != lexer.Tok_eof; {
		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			return
		}

		token.PrintToken()

		token, err = lex.GetToken()
	}
}
