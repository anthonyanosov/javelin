package main

import (
	"flag"
	"fmt"
	"javelin/lexer"
	"javelin/parser"
	"os"
)

func main() {
	// Input reading
	src := flag.String("src", "main.jv", "Javelin source file to compile")
	flag.Parse()

	text, err := os.ReadFile(*src)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open file:", err)
		os.Exit(1)
	}

	// Lexing
	lex := lexer.NewLexer(string(text))
	lex.Tokenize()

	// Parsing
	parser := parser.NewParser(lex.Tokens)
	program := parser.Parse()
	fmt.Println(program)
	os.Exit(0)
}
