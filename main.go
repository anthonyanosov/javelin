package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
)

func main() {
	// Take in an arg - one Go file
	src := flag.String("src", "main.go", "Go source file to analyze")
	flag.Parse()

	// Search for the file
	path, err := filepath.Abs(*src)
	if err != nil {
		fmt.Println("javelin could not find the Go file requested for analysis")
		return
	}
	astFile, err := parseFile(path)
	if err != nil {
		fmt.Println("javelin encountered an error while parsing the Go file")
	}

	// Look through the AST

}

func parseFile(path string) (*ast.File, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return file, nil
}
