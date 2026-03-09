package main

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/sofk69/loglint/pkg/analyzer"
)

func init() {
	register.Plugin("loglint", analyzer.Analyzer)
}
