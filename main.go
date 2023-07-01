// callgraph-viewer: a tool for visualizing the call graph of a Go program.
// See Usage for details, or run with -help.
package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/xat0mz/go-callgraph-viewer/pkg/analyzer"
	"github.com/xat0mz/go-callgraph-viewer/pkg/flags"
)

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

func main() {
	f := new(flags.Flags)
	f.Parse()

	a := new(analyzer.Analyzer)
	if err := a.DoAnalyze(f); err != nil {
		fmt.Fprintf(os.Stderr, "callgraph-viewer: %s\n", err)
		os.Exit(1)
	}
}
