package analyzer

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"

	"golang.org/x/tools/go/callgraph"
)

func (a *Analyzer) GenerateGraphFor(pkg string, fn string) (error, error) {
	for f, n := range a.cg.Nodes {
		// rel := f.RelString(nil)
		if strings.Contains(f.Pkg.Pkg.Path(), pkg) {
			fmt.Println(
				f.Pkg.Pkg.Path(),
				n.Func.Name(),
				f.Pkg.Members,
			)
		}
	}
	fmt.Println()
	// if err := callgraph.GraphVisitEdges(a.cg, func(edge *callgraph.Edge) error {
	// 	return nil
	// }); err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

// -- output------------------------------------------------------------
func (a *Analyzer) OutputGraph(format string) error {
	var stdout io.Writer = os.Stdout
	var before, after string

	prog := a.prog
	cg := a.cg

	// Pre-canned formats.
	switch format {
	case "digraph":
		format = `{{printf "%q %q" .Caller .Callee}}`

	case "graphviz":
		before = "digraph callgraph {\n"
		after = "}\n"
		format = `  {{printf "%q" .Caller}} -> {{printf "%q" .Callee}}`
	}

	tmpl, err := template.New("-format").Parse(format)
	if err != nil {
		return fmt.Errorf("invalid -format template: %v", err)
	}

	// Allocate these once, outside the traversal.
	var buf bytes.Buffer
	data := Edge{fset: prog.Fset}

	fmt.Fprint(stdout, before)
	if err := callgraph.GraphVisitEdges(cg, func(edge *callgraph.Edge) error {
		data.position.Offset = -1
		data.edge = edge
		data.Caller = edge.Caller.Func
		data.Callee = edge.Callee.Func

		buf.Reset()
		if err := tmpl.Execute(&buf, &data); err != nil {
			return err
		}
		stdout.Write(buf.Bytes())
		if len := buf.Len(); len == 0 || buf.Bytes()[len-1] != '\n' {
			fmt.Fprintln(stdout)
		}
		return nil
	}); err != nil {
		return err
	}
	fmt.Fprint(stdout, after)
	return nil
}
