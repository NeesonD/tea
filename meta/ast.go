package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func AstExample() {

	var ruleCtx RuleContext
	fmt.Println(ruleCtx)

	config := ` isNewStudent(ruleCtx) && (hasEnoughCurriculum(ruleCtx) || inWhiteList(ruleCtx))`

	fileSet := token.NewFileSet()
	expr, err := parser.ParseExpr(config)
	if err != nil {
		log.Fatal(err)
	}

	ast.Print(fileSet, expr)

	ast.Inspect(expr, func(node ast.Node) bool {

		switch node.(type) {
		case *ast.CallExpr:
			callExpr := node.(*ast.CallExpr)
			fmt.Println(callExpr)
		}

		return true
	})
}

var RuleDict = map[string]rule{
	"isNewStudent":        isNewStudent,
	"hasEnoughCurriculum": hasEnoughCurriculum,
	"inWhiteList":         inWhiteList,
}

type RuleContext struct {
}

type rule func(ruleContext RuleContext) bool

func isNewStudent(ruleCtx RuleContext) bool {
	return false
}

func hasEnoughCurriculum(ruleCtx RuleContext) bool {
	return false
}

func inWhiteList(ruleCtx RuleContext) bool {
	return false
}
