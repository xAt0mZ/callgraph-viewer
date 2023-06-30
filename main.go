// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// go-callgraph-viewer: a tool for visualizing the call graph of a Go program.
// See Usage for details, or run with -help.
package main // import "golang.org/x/tools/cmd/callgraph"

import (
	"flag"
	"fmt"
	"go/build"
	"os"
	"runtime"

	"github.com/xat0mz/go-callgraph-viewer/pkg/analyzer"
	"golang.org/x/tools/go/buildutil"
)

// flags
var (
	algoFlag = flag.String("algo", "rta",
		`Call graph construction algorithm (cha, rta, vta).`)
	testFlag = flag.Bool("test", false,
		`Loads test code (*_test.go) for imported packages.`)
	urlFlag = flag.String("url", ":8080",
		`An ip:port for the server`)

	// 	formatFlag = flag.String("format",
	// 		"{{.Caller}}\t--{{.Dynamic}}-{{.Line}}:{{.Column}}-->\t{{.Callee}}",
	// 		`A template expression specifying how to format an edge.
	// 		default: {{.Caller}}\t--{{.Dynamic}}-{{.Line}}:{{.Column}}-->\t{{.Callee}}
	// `)
)

func init() {
	flag.Var((*buildutil.TagsFlag)(&build.Default.BuildTags), "tags", buildutil.TagsFlagDoc)
}

func init() {
	// If $GOMAXPROCS isn't set, use the full capacity of the machine.
	// For small machines, use at least 4 threads.
	if os.Getenv("GOMAXPROCS") == "" {
		n := runtime.NumCPU()
		if n < 4 {
			n = 4
		}
		runtime.GOMAXPROCS(n)
	}
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "callgraph: %s\n", err)
	os.Exit(1)
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Fprint(os.Stderr, Usage)
		return
	}

	a := analyzer.Analyzer{}

	if err := a.BuildProgram("", "", *testFlag, flag.Args()); err != nil {
		exit(err)
	}

	if err := a.BuildGraph(*algoFlag); err != nil {
		exit(err)
	}

	if err := a.Serve(*urlFlag); err != nil {
		exit(err)
	}

	// if err := analyzer.OutputGraph(*formatFlag); err != nil {
	// 	exit(err)
	// }

}

const Usage = `go-callgraph-viewer: display the call graph of a Go program in a nice UI.

Usage:

  go-callgraph-viewer [-algo=cha|rta|vta] [-test] package...

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
