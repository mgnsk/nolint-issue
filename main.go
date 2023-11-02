package main

import (
	"github.com/kisielk/errcheck/errcheck"
	"github.com/kyoh86/nolint"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	var checkers []*analysis.Analyzer

	checkers = append(checkers, nolint.WrapFunc(errcheck.Analyzer, func(analysis.Diagnostic) bool {
		return false
	}))

	multichecker.Main(checkers...)
}
