package analyzer

import (
	"github.com/xat0mz/callgraph-viewer/pkg/flags"
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/ssa"
)

type Analyzer struct {
	flags *flags.Flags
	prog  *ssa.Program
	pkgs  []*ssa.Package
	cg    *callgraph.Graph

	nodes map[string]*callgraph.Node
}

func (a *Analyzer) Analyze(f *flags.Flags) error {
	a.flags = f

	if err := a.buildProgram(); err != nil {
		return err
	}

	if err := a.buildGraph(); err != nil {
		return err
	}

	a.buildNodesMap()

	return nil
}
