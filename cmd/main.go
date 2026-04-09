package main

import (
	"flag"
	"fmt"
	"github.com/anthonyanosov/javelin/pkg"
)

func main() {
	src := flag.String("src", "main.go", "Go source file to analyze")
	flag.Parse()

	funs, err := javelin.AnalyzeFile(*src)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, f := range funs {
		fmt.Printf("%+v\n", f)
	}
}
