package server

import (
	"fmt"
	"net/http"

	"github.com/xat0mz/callgraph-viewer/pkg/analyzer"
	"github.com/xat0mz/callgraph-viewer/pkg/flags"
)

type Server struct {
	url      string
	analyzer *analyzer.Analyzer
}

func NewServer(f *flags.Flags, a *analyzer.Analyzer) *Server {
	server := new(Server)
	server.analyzer = a
	server.url = f.Url
	return server
}

func (s *Server) Serve() error {
	url := s.url

	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)

	http.HandleFunc("/api", s.handleAPI)

	fmt.Printf("Listening on %q ...\n", url)
	return http.ListenAndServe(url, nil)
}
