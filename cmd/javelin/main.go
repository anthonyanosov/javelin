package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/anthonyanosov/javelin/pkg"
	"os"
)

func main() {
	src := flag.String("src", "main.go", "Go source file to analyze")
	flag.Parse()

	funs, err := javelin.AnalyzeFile(*src)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(funs); err != nil {
		fmt.Fprintln(os.Stderr, "Error encoding JSON:", err)
		os.Exit(1)
	}
}
