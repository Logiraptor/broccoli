package main

import (
	"bufio"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	fs := token.NewFileSet()
	rd := bufio.NewScanner(os.Stdin)
	for rd.Scan() {
		expr, err := parser.ParseExpr(rd.Text())
		if err != nil {
			log.Fatal(err.Error())
		}

		err = ast.Fprint(os.Stdin, fs, expr, nil)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
