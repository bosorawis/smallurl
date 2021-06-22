package server

import (
	"encoding/json"
	"fmt"
	"github.com/dihmuzikien/smallurl/domain/url"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	db url.Repository
	router chi.Router
}

func NewServer(r url.Repository) (*Server, error){
	router := chi.NewRouter()
	server := &Server{
		db: r,
		router: router,
	}
	server.routes()
	return server, nil
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int){
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) decode(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}


func (s *Server) handleListUrl() http.HandlerFunc {
	type model struct {
		ID          string `json:"id"`
		Destination string `json:"destination"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := s.db.List(r.Context())
		if err != nil {
			fmt.Println("failed to get url", err)
			http.Error(w, "failed to list URL", 500)
			return
		}
		m := make([]model, len(items))
		for i, v := range items {
			m[i] = model{
				ID:          v.ID,
				Destination: v.Destination,
			}
		}
		data, err := json.Marshal(items)
		if err != nil {
			fmt.Println("failed to parse response", err)
			http.Error(w, "failed to list URL", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}


func (s *Server) handleCreateUrl() http.HandlerFunc {
	type request struct {
		ID          string `json:"id"`
		Destination string `json:"destination"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var d request
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			fmt.Println("failed to parse body", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = s.db.Put(r.Context(),d.ID, d.Destination)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}