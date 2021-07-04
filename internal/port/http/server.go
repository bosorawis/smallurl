package http

import (
	"errors"
	"github.com/dihmuzikien/smallurl/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	svc    domain.UrlUseCase
	router *gin.Engine
}

func New(svc domain.UrlUseCase) (*Server, error) {
	router := gin.Default()
	server := &Server{
		svc:    svc,
		router: router,
	}
	server.routes()
	return server, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) handleListUrl() gin.HandlerFunc {
	type response struct {
		ID          string `json:"id"`
		Destination string `json:"destination"`
	}
	return func(c *gin.Context) {
		items, err := s.svc.List(c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		responseData := make([]response, len(items))
		for i, v := range items {
			responseData[i] = response{
				ID:          v.ID,
				Destination: v.Destination,
			}
		}
		c.JSON(http.StatusOK, responseData)
	}
}

func (s *Server) handleCreateUrl() gin.HandlerFunc {
	type request struct {
		Destination string `json:"destination"`
	}
	type response struct {
		ID string `json:"id"`
	}
	return func(c *gin.Context) {
		var d request
		c.BindJSON(&d)
		url, err := s.svc.Create(c, d.Destination)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		resp := response{
			ID: url.ID,
		}
		c.JSON(http.StatusCreated, resp)
	}
}

func (s *Server) handleCreateUrlWithAlias() gin.HandlerFunc {
	type request struct {
		Alias       string `json:"alias"`
		Destination string `json:"destination"`
	}
	type response struct {
		ID string `json:"id"`
	}
	return func(c *gin.Context) {
		var req request
		c.BindJSON(&req)
		if len(req.Alias) <= 3 {
			c.AbortWithError(http.StatusBadRequest, errors.New("alias must be longer than 3 characters"))
			return
		}

		url, err := s.svc.CreateWithAlias(c, req.Alias, req.Destination)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		resp := response{
			ID: url.ID,
		}
		c.JSON(http.StatusCreated, resp)
	}
}
