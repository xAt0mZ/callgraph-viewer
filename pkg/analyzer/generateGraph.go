package analyzer

import "golang.org/x/tools/go/callgraph"

func (a *Analyzer) generateGraph(pkg string, fn string) (error, error) {
	if err := callgraph.GraphVisitEdges(a.cg, func(edge *callgraph.Edge) error {
		return nil
	}); err != nil {
		return nil, err
	}
	return nil, nil
}
