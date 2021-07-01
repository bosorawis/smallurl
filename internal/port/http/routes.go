package http

func (s *Server) routes() {
	s.router.MethodFunc("POST", "/v1", s.handleCreateUrl())
	s.router.MethodFunc("POST", "/v1/alias", s.handleCreateUrlWithAlias())
	s.router.MethodFunc("GET", "/v1", s.handleListUrl())
}