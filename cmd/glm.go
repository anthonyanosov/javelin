package main

import (
	"flag"
	"fmt"
	"golem/lexer"
	"os"
)

func main() {
	src := flag.String("src", "main.gol", "Golem source file to compile")
	flag.Parse()

	text, err := os.ReadFile(*src)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open file:", err)
		os.Exit(1)
	}

	lex := lexer.NewLexer(string(text))

	lex.Tokenize()
	lex.PrintTokens()
	os.Exit(0)
}
