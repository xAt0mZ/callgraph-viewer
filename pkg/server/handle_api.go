package server

import (
	"encoding/json"
	"net/http"

	"golang.org/x/tools/go/callgraph"
)

type Response struct {
	nodes []*callgraph.Node ``
}

func (s *Server) handleAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pkg := r.URL.Query().Get("pkg")
	fn := r.URL.Query().Get("fn")

	_, err := s.analyzer.GenerateGraphFor(pkg, fn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	res := new(Response)
	res.nodes = nil

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
