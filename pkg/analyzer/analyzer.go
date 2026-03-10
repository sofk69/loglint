package analyzer

import (
	"fmt"
	"go/ast"

	"github.com/sofk69/loglint/pkg/analyzer/rules"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "checks log messages",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {

	for _, file := range pass.Files {

		ast.Inspect(file, func(n ast.Node) bool {

			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if !isLogCall(call) {
				return true
			}

			if len(call.Args) == 0 {
				return true
			}

			msg, ok := extractMessage(call.Args[0])
			if !ok {
				return true
			}

			validate(pass, call.Args[0], msg)

			return true
		})

	}

	return nil, nil
}

func validate(pass *analysis.Pass, node ast.Node, msg string) {
	diagnostics := []string{}

	if !rules.IsLowercase(msg) {
		diagnostics = append(diagnostics, "lowercase")
		pass.Reportf(node.Pos(), "log message should start with lowercase")
	}

	if !rules.IsEnglish(msg) {
		diagnostics = append(diagnostics, "english")
		pass.Reportf(node.Pos(), "log message must be english")
	}

	if rules.HasSpecialChars(msg) {
		diagnostics = append(diagnostics, "special")
		pass.Reportf(node.Pos(), "log message contains special characters")
	}

	if rules.ContainsSensitive(msg) {
		diagnostics = append(diagnostics, "sensitive")
		pass.Reportf(node.Pos(), "log message may contain sensitive data")
	}

	if len(diagnostics) > 0 {
		fmt.Printf("Line with message %q, Diagnostics: %v\n", msg, diagnostics)
	}
}
