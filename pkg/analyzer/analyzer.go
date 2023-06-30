package analyzer

import (
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/ssa"
)

//////////////
// ANALYZER //
//////////////

type Analyzer struct {
	prog *ssa.Program
	pkgs []*ssa.Package
	cg   *callgraph.Graph
}
