package analyzer

import "golang.org/x/tools/go/callgraph"

func (a *Analyzer) buildNodesMap() {
	a.nodes = make(map[string]*callgraph.Node)

	for f, n := range a.cg.Nodes {
		a.nodes[f.String()] = n
	}
}
