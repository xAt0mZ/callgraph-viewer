// callgraph-viewer: a tool for visualizing the call graph of a Go program.
// See Usage for details, or run with -help.
package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/xat0mz/callgraph-viewer/pkg/analyzer"
	"github.com/xat0mz/callgraph-viewer/pkg/flags"
	"github.com/xat0mz/callgraph-viewer/pkg/server"
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

func exit(err error) {
	fmt.Fprintf(os.Stderr, "callgraph-viewer: %s\n", err)
	os.Exit(1)
}

func main() {
	f := new(flags.Flags)
	f.Parse()

	a := new(analyzer.Analyzer)
	if err := a.Analyze(f); err != nil {
		exit(err)
	}

	s := server.NewServer(f, a)
	if err := s.Serve(); err != nil {
		exit(err)
	}
}
