//go:build golangci

package main

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/sofk69/loglint/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("loglint", New)
}

func New() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}, nil
}
