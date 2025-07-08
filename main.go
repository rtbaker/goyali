package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/rtbaker/goyali/lexer"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(strings.NewReader("    abcdefg abc(dfdf)/2+"))
	lex := lexer.NewLexer(reader)

	fmt.Println("Start typing: ")

	var token *lexer.Token
	var err error

	for token, err = lex.GetToken(); err == nil && token.Code != lexer.EOF; {
		token.PrintToken()

		// Get next
		token, err = lex.GetToken()
	}

	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
}
