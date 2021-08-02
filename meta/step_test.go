package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"strconv"
	"testing"
)

// 拼接
func TestGenCode(t *testing.T) {
	//createFile(genFileName("step1"), step1Code, false)
	createFileAndFormat(genFileName("step1"), step1Code)
}

// template
func TestGenCode2(t *testing.T) {

	templateInfo := &TemplateInfo{
		Text: ExampleText,
		data: map[string]interface{}{
			"Biz": "Group",
		},
	}
	buffer, err := templateInfo.GenCodeByTemplate()
	if err != nil {
		log.Fatal(err)
	}
	createFileAndFormat(genFileName("step2"), buffer.String())
}

// ast
func TestGenCode3(t *testing.T) {

	fset := token.NewFileSet()
	path, _ := filepath.Abs("./gen/gen_step2.go")
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		panic(err)
	}
	ast.Print(fset, f)

	ast.Inspect(f, func(node ast.Node) bool {
		switch node.(type) {
		case *ast.GenDecl:
			genDecl := node.(*ast.GenDecl)
			if genDecl.Tok == token.IMPORT {
				hasImported := false
				for _, spec := range genDecl.Specs {
					importSpec := spec.(*ast.ImportSpec)
					if importSpec.Path.Value == strconv.Quote("context") {
						hasImported = true
					}
				}
				// 没有context，就加入context
				if !hasImported {
					genDecl.Specs = append(genDecl.Specs, &ast.ImportSpec{Path: &ast.BasicLit{
						Kind:  token.STRING,
						Value: strconv.Quote("context"),
					}})
				}
			}
		case *ast.FuncDecl:
			funcDecl := node.(*ast.FuncDecl)
			ft := funcDecl.Type
			hasContext := false
			for _, v := range ft.Params.List {
				if expr, ok := v.Type.(*ast.SelectorExpr); ok {
					if ident, ok := expr.X.(*ast.Ident); ok {
						if ident.Name == "context" {
							hasContext = true
						}
					}
				}
			}
			if !hasContext {
				ctxFiled := &ast.Field{
					Names: []*ast.Ident{ast.NewIdent("ctx")},
					Type: &ast.SelectorExpr{
						X:   ast.NewIdent("context"),
						Sel: ast.NewIdent("Context"),
					},
				}
				list := []*ast.Field{ctxFiled}
				ft.Params.List = append(list, ft.Params.List...)
			}
			return false
		}
		return true
	})
	var output []byte
	buffer := bytes.NewBuffer(output)
	err = format.Node(buffer, fset, f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buffer)
	createFileAndFormat(genFileName("step3"), buffer.String())
}
