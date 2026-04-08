package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
)

type FunctionMetric struct {
	Name       string
	Complexity int
}

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
	var funs []FunctionMetric
	ast.Inspect(astFile, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			complexity := 1

			ast.Inspect(fn.Body, func(n ast.Node) bool {
				// For nested anonymous functions, step over it
				if _, ok := n.(*ast.FuncDecl); ok {
					return false
				}
				switch t := n.(type) {
				case *ast.IfStmt:
					complexity++
				case *ast.ForStmt:
					complexity++
				case *ast.RangeStmt:
					complexity++
				case *ast.BinaryExpr:
					if t.Op == token.LAND || t.Op == token.LOR {
						complexity++
					}
				}
				return true
			})

			funs = append(funs, FunctionMetric{
				Name:       fn.Name.Name,
				Complexity: complexity,
			})

			return false
		}
		return true
	})

	for _, f := range funs {
		fmt.Printf("%+v\n", f)
	}
}

func parseFile(path string) (*ast.File, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return file, nil
}
