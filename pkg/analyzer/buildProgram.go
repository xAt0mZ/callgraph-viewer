package analyzer

import (
	"fmt"
	"os"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

var PackagesLoadMode = packages.NeedDeps |
	packages.NeedSyntax |
	packages.NeedTypesInfo |
	packages.NeedTypes |
	packages.NeedTypesSizes |
	packages.NeedImports |
	packages.NeedName |
	packages.NeedFiles |
	packages.NeedCompiledGoFiles

// -- program construction ------------------------------------------
func (a *Analyzer) BuildProgram(dir, gopath string, tests bool, args []string) error {
	cfg := &packages.Config{
		Mode:  PackagesLoadMode,
		Tests: tests,
		Dir:   dir,
	}
	if gopath != "" {
		cfg.Env = append(os.Environ(), "GOPATH="+gopath) // to enable testing
	}
	initial, err := packages.Load(cfg, args...)
	if err != nil {
		return err
	}
	if packages.PrintErrors(initial) > 0 {
		return fmt.Errorf("packages contain errors")
	}

	// Create and build SSA-form program representation.
	mode := ssa.InstantiateGenerics // instantiate generics by default for soundness
	prog, pkgs := ssautil.AllPackages(initial, mode)
	prog.Build()

	a.prog = prog
	a.pkgs = pkgs
	return nil
}
