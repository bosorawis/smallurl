package http

import (
	"encoding/json"
	"fmt"
	"github.com/dihmuzikien/smallurl/domain"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	svc domain.UrlUseCase
	router chi.Router
}

func NewServer(svc domain.UrlUseCase) (*Server, error){
	router := chi.NewRouter()
	server := &Server{
		svc: svc,
		router: router,
	}
	server.routes()
	return server, nil
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int){
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
}

func (s *Server) decode(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}


func (s *Server) handleListUrl() http.HandlerFunc {
	type response struct {
		ID          string `json:"id"`
		Destination string `json:"destination"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := s.svc.List(r.Context())
		if err != nil {
			fmt.Println("failed to get url", err)
			http.Error(w, "failed to list URL", 500)
			return
		}
		m := make([]response, len(items))
		for i, v := range items {
			m[i] = response{
				ID:          v.ID,
				Destination: v.Destination,
			}
		}
		s.respond(w, r, items, http.StatusOK)
	}
}


func (s *Server) handleCreateUrl() http.HandlerFunc {
	type request struct {
		Destination string `json:"destination"`
	}
	type response struct {
		ID string `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var d request
		err := s.decode(w, r, &d)
		if err != nil {
			fmt.Println("failed to parse body", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		url, err := s.svc.Create(r.Context(), d.Destination)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp := response{
			ID: url.ID,
		}
		s.respond(w, r, resp, http.StatusCreated)
	}
}