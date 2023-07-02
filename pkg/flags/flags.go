package flags

import (
	"flag"
	"fmt"
	"go/build"
	"os"

	"golang.org/x/tools/go/buildutil"
)

type Flags struct {
	Dir    string
	Gopath string
	Algo   string
	Tests  bool
	Url    string
	Args   []string
}

// flags
// var (
// 	formatFlag = flag.String("format",
// 		"{{.Caller}}\t--{{.Dynamic}}-{{.Line}}:{{.Column}}-->\t{{.Callee}}",
// 		`A template expression specifying how to format an edge.
// `)
// )

func (f *Flags) Parse() {
	f.Dir = ""
	f.Gopath = ""
	flag.StringVar(&f.Algo, "algo", "rta", "Call graph construction algorithm (cha, rta, vta)")
	flag.BoolVar(&f.Tests, "test", false, "Loads test code (*_test.go) for imported packages")
	flag.StringVar(&f.Url, "url", ":8080", "ip:port for the server")

	flag.Var((*buildutil.TagsFlag)(&build.Default.BuildTags), "tags", buildutil.TagsFlagDoc)

	flag.Parse()

	f.Args = flag.Args()
	if len(f.Args) == 0 {
		fmt.Fprint(os.Stderr, Usage)
		os.Exit(0)
	}
}

const Usage = `callgraph-viewer: display the call graph of a Go program in a nice UI.

Usage:

  callgraph-viewer [-algo=cha|rta|vta] [-test] [-port=":8080"] package...

Flags:

-algo      Specifies the call-graph construction algorithm, one of:

            cha         Class Hierarchy Analysis
            rta         Rapid Type Analysis (default)
            vta         Variable Type Analysis

          The algorithms are ordered by increasing precision in their
          treatment of dynamic calls (and thus also computational cost).
          RTA requires a whole program (main or test), and
          include only functions reachable from main.

-test     Include the package's tests in the analysis.

-port			Server ip:port (default ":8080" )

It starts an http server on port 8080 that allows to query the graph
	- Opening a browser on localhost:8080/ gives a UI to display the graph
	- calls with query params allow to query the graph (callers and called) functions

  	Example: curl localhost:8080/?pkg="github.com/xat0mz/go-callgraph-viewer"&fn="main"
`

// -format    Specifies the format in which each call graph edge is displayed.
//            One of:

//             digraph     output suitable for input to
//                         golang.org/x/tools/cmd/digraph.
//             graphviz    output in AT&T GraphViz (.dot) format.

//            All other values are interpreted using text/template syntax.
//            The default value is:

//             {{.Caller}}\t--{{.Dynamic}}-{{.Line}}:{{.Column}}-->\t{{.Callee}}

//            The structure passed to the template is (effectively):

//                    type Edge struct {
//                            Caller      *ssa.Function // calling function
//                            Callee      *ssa.Function // called function

//                            // Call site:
//                            Filename    string // containing file
//                            Offset      int    // offset within file of '('
//                            Line        int    // line number
//                            Column      int    // column number of call
//                            Dynamic     string // "static" or "dynamic"
//                            Description string // e.g. "static method call"
//                    }

//            Caller and Callee are *ssa.Function values, which print as
//            "(*sync/atomic.Mutex).Lock", but other attributes may be
//            derived from them, e.g. Caller.Pkg.Pkg.Path yields the
//            import path of the enclosing package.  Consult the go/ssa
//            API documentation for details.
