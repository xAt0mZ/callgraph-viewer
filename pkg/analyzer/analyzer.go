package analyzer

import (
	"github.com/xat0mz/go-callgraph-viewer/pkg/flags"
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/ssa"
)

type Analyzer struct {
	flags *flags.Flags
	prog  *ssa.Program
	pkgs  []*ssa.Package
	cg    *callgraph.Graph
}

func (a *Analyzer) DoAnalyze(f *flags.Flags) error {
	a.flags = f

	if err := a.buildProgram(); err != nil {
		return err
	}

	if err := a.buildGraph(); err != nil {
		return err
	}

	return a.serve()
	// if err := analyzer.OutputGraph(*formatFlag); err != nil {
	// 	exit(err)
	// }

}
