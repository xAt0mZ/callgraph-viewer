package digraph

import "fmt"

type Nodelist []string

func (l Nodelist) println(sep string) {
	for i, node := range l {
		if i > 0 {
			fmt.Fprint(stdout, sep)
		}
		fmt.Fprint(stdout, node)
	}
	fmt.Fprintln(stdout)
}
