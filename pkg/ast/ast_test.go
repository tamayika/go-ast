package ast

import (
	"encoding/json"
	"fmt"
	"testing"
)

var src = `package main

import "fmt"

func main() {
	fmt.Printf("Hello world")
}
`
var expected = `{"type":"ast.File","pos":1,"end":71,"children":[{"type":"ast.Ident","pos":9,"end":13,"children":[]},{"type":"ast.GenDecl","pos":15,"end":27,"children":[{"type":"ast.ImportSpec","pos":22,"end":27,"children":[{"type":"ast.BasicLit","pos":22,"end":27,"children":[]}]}]},{"type":"ast.FuncDecl","pos":29,"end":71,"children":[{"type":"ast.Ident","pos":34,"end":38,"children":[]},{"type":"ast.FuncType","pos":29,"end":40,"children":[{"type":"ast.FieldList","pos":38,"end":40,"children":[]}]},{"type":"ast.BlockStmt","pos":41,"end":71,"children":[{"type":"ast.ExprStmt","pos":44,"end":69,"children":[{"type":"ast.CallExpr","pos":44,"end":69,"children":[{"type":"ast.SelectorExpr","pos":44,"end":54,"children":[{"type":"ast.Ident","pos":44,"end":47,"children":[]},{"type":"ast.Ident","pos":48,"end":54,"children":[]}]},{"type":"ast.BasicLit","pos":55,"end":68,"children":[]}]}]}]}]}]}`

func TestParse(t *testing.T) {
	n, err := Parse(src)
	if err != nil {
		t.Fatal(err)
	}
	json, err := json.Marshal(n)
	if err != nil {
		t.Fatal(err)
	}
	if string(json) != expected {
		fmt.Println(string(json))
		fmt.Println(expected)
		t.Error("Not matched")
	}
}
