package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Sinubio/mycalcservice/pkg/calculate"
)

type Server struct {
	addr string
}

func New() *Server {
	return &Server{
		addr: ":8080",
	}
}

func (s *Server) RunServer() {
	http.HandleFunc("/api/v1/calculate", s.handleCalculate)

	log.Printf("Starting server on %s...\n", s.addr)
	if err := http.ListenAndServe(s.addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (s *Server) handleCalculate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Expression string `json:"expression"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result, err := calculate.Calc(req.Expression)
	if err != nil {
		http.Error(w, "Expression is not valid", http.StatusUnprocessableEntity)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": result,
	})
}
