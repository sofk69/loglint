package analyzer

import (
	"go/ast"
	"go/token"
	"strconv"
)

func isLogCall(call *ast.CallExpr) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	if pkg, ok := sel.X.(*ast.Ident); ok {
		if pkg.Name == "slog" {
			switch sel.Sel.Name {
			case "Info", "Error", "Warn", "Debug":
				return true
			}
		}
		if pkg.Name == "zap" {
			switch sel.Sel.Name {
			case "Info", "Error", "Warn", "Debug":
				return true
			}
		}
	}
	return false
}

func extractMessage(expr ast.Expr) (string, bool) {

	switch v := expr.(type) {

	case *ast.BasicLit:

		if v.Kind == token.STRING {
			s, err := strconv.Unquote(v.Value)
			if err != nil {
				return "", false
			}
			return s, true
		}

	case *ast.BinaryExpr:

		left, ok := extractMessage(v.X)
		if ok {
			return left, true
		}

	case *ast.CallExpr:

		if isFmtSprintf(v) {
			if len(v.Args) == 0 {
				return "", false
			}
			return extractMessage(v.Args[0])
		}

	}

	return "", false
}

func isFmtSprintf(call *ast.CallExpr) bool {

	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	pkg, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}

	return pkg.Name == "fmt" && sel.Sel.Name == "Sprintf"
}
