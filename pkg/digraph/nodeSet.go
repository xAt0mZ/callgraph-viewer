package digraph

import "sort"

type nodeset map[string]bool

func (s nodeset) sort() Nodelist {
	nodes := make(Nodelist, len(s))
	var i int
	for node := range s {
		nodes[i] = node
		i++
	}
	sort.Strings(nodes)
	return nodes
}

func (s nodeset) addAll(x nodeset) {
	for node := range x {
		s[node] = true
	}
}
