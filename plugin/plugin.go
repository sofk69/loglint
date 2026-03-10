//go:build golangci

package main

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/mitchellh/mapstructure"
	"github.com/sofk69/loglint/pkg/analyzer"
	"github.com/sofk69/loglint/pkg/analyzer/config"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("loglint", New)
}

func New(settings interface{}) ([]*analysis.Analyzer, error) {
	var cfg config.Config
	if settings != nil {
		if err := mapstructure.Decode(settings, &cfg); err != nil {
			return nil, err
		}
	}
	analyzer.SetConfig(cfg)
	return []*analysis.Analyzer{analyzer.Analyzer}, nil
}
