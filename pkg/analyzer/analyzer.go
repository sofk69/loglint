package analyzer

import (
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

	if !rules.IsLowercase(msg) {
		pass.Reportf(node.Pos(), "log message should start with lowercase")
	}

	if !rules.IsEnglish(msg) {
		pass.Reportf(node.Pos(), "log message must be english")
	}

	if rules.HasSpecialChars(msg) {
		pass.Reportf(node.Pos(), "log message contains special characters")
	}

	if rules.ContainsSensitive(msg) {
		pass.Reportf(node.Pos(), "log message may contain sensitive data")
	}
}
