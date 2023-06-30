package analyzer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"

	"golang.org/x/tools/go/callgraph"
)

// -- serve http ------------------------------------------------------------
func (a *Analyzer) Serve(url string) error {
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		pkg := r.URL.Query().Get("pkg")
		fn := r.URL.Query().Get("fn")

		_, err := a.generateGraph(pkg, fn)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(nil)
	})

	fmt.Printf("Listening on %q ...", url)
	return http.ListenAndServe(url, nil)
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
