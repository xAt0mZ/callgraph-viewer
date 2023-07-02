package analyzer

import (
	"fmt"

	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/callgraph/cha"
	"golang.org/x/tools/go/callgraph/rta"
	"golang.org/x/tools/go/callgraph/vta"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// -- callgraph construction ------------------------------------------
func (a *Analyzer) buildGraph() error {
	prog := a.prog
	pkgs := a.pkgs
	algo := a.flags.Algo

	var cg *callgraph.Graph

	switch algo {
	case "cha":
		cg = cha.CallGraph(prog)

	case "rta":
		mains, err := mainPackages(pkgs)
		if err != nil {
			return err
		}
		var roots []*ssa.Function
		for _, main := range mains {
			roots = append(roots, main.Func("init"), main.Func("main"))
		}
		rtares := rta.Analyze(roots, true)
		cg = rtares.CallGraph

		// NB: RTA gives us Reachable and RuntimeTypes too.

	case "vta":
		cg = vta.CallGraph(ssautil.AllFunctions(prog), cha.CallGraph(prog))

	default:
		return fmt.Errorf("unknown algorithm: %s", algo)
	}

	cg.DeleteSyntheticNodes()

	a.cg = cg
	return nil
}

// mainPackages returns the main packages to analyze.
// Each resulting package is named "main" and has a main function.
func mainPackages(pkgs []*ssa.Package) ([]*ssa.Package, error) {
	var mains []*ssa.Package
	for _, p := range pkgs {
		if p != nil && p.Pkg.Name() == "main" && p.Func("main") != nil {
			mains = append(mains, p)
		}
	}
	if len(mains) == 0 {
		return nil, fmt.Errorf("no main packages")
	}
	return mains, nil
}
