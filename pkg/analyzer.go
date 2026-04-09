package javelin

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func ParseFile(path string) (*ast.File, *token.FileSet, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, err
	}
	return file, fset, nil
}

func ComputeComplexity(fn *ast.FuncDecl) int {
	complexity := 1
	ast.Inspect(fn.Body, func(n ast.Node) bool {
		if _, ok := n.(*ast.FuncDecl); ok {
			return false
		}
		switch t := n.(type) {
		case *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.CaseClause:
			complexity++
		case *ast.BinaryExpr:
			if t.Op == token.LAND || t.Op == token.LOR {
				complexity++
			}
		}
		return true
	})
	return complexity
}

func AnalyzeFile(path string) ([]FunctionMetric, error) {
	astFile, fset, err := ParseFile(path)
	if err != nil {
		return nil, err
	}

	var funs []FunctionMetric
	ast.Inspect(astFile, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			pos := fset.Position(fn.Pos())
			end := fset.Position(fn.End())
			funs = append(funs, FunctionMetric{
				Name:       fn.Name.Name,
				Complexity: ComputeComplexity(fn),
				StartLine:  pos.Line - 1,
				EndLine:    end.Line - 1,
			})
			return false
		}
		return true
	})

	return funs, nil
}
