package digraph

// A graph maps nodes to the non-nil set of their immediate successors.
type graph map[string]nodeset

func (g graph) addNode(node string) nodeset {
	edges := g[node]
	if edges == nil {
		edges = make(nodeset)
		g[node] = edges
	}
	return edges
}

func (g graph) addEdges(from string, to ...string) {
	edges := g.addNode(from)
	for _, to := range to {
		g.addNode(to)
		edges[to] = true
	}
}
