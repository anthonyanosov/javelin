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
	// outFile := flag.String("o", "a.out", "Output executable file name")
	flag.Parse()

	if len(flag.Args()[0]) < 1 {
		panic("Usage: jav <source_file.jv> [-o output_file]")
	}

	src := flag.Args()[0]
	text, err := os.ReadFile(src)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open file:", err)
		os.Exit(1)
	}

	// Lexing
	lex := lexer.NewLexer(string(text))
	lex.Tokenize()

	// Parsing
	parser := parser.NewParser(lex.Tokens)
	parser.Parse()
	os.Exit(0)
}
