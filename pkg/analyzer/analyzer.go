package analyzer

import (
	"flag"
	"go/ast"

	"github.com/sofk69/loglint/pkg/analyzer/config"
	"github.com/sofk69/loglint/pkg/analyzer/rules"
	"golang.org/x/tools/go/analysis"
)

var cfg config.Config

var Analyzer = &analysis.Analyzer{
	Name:  "loglint",
	Doc:   "checks log messages",
	Run:   run,
	Flags: flags(),
}

func flags() flag.FlagSet {
	fs := flag.NewFlagSet("loglint", flag.ContinueOnError)
	fs.BoolVar(&cfg.DisableLowercase, "disable-lowercase", false, "disable lowercase check")
	fs.BoolVar(&cfg.DisableEnglish, "disable-english", false, "disable english check")
	fs.BoolVar(&cfg.DisableSpecial, "disable-special-chars", false, "disable special chars check")
	fs.BoolVar(&cfg.DisableSensitive, "disable-sensitive", false, "disable sensitive data check")
	return *fs
}

func SetConfig(c config.Config) {
	cfg = c
	rules.SetExtraSensitive(c.ExtraSensitive)
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
	if !cfg.DisableLowercase && !rules.IsLowercase(msg) {
		pass.Reportf(node.Pos(), "log message should start with lowercase")
	}

	if !cfg.DisableEnglish && !rules.IsEnglish(msg) {
		pass.Reportf(node.Pos(), "log message must be english")
	}

	if !cfg.DisableSpecial && rules.HasSpecialChars(msg) {
		pass.Reportf(node.Pos(), "log message contains special characters")
	}

	if !cfg.DisableSensitive && rules.ContainsSensitive(msg) {
		pass.Reportf(node.Pos(), "log message may contain sensitive data")
	}
}
