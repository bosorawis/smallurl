package http

func (s *Server) routes() {
	s.router.POST("/v1", s.handleCreateUrl())
	s.router.POST("/v1/alias", s.handleCreateUrlWithAlias())
	s.router.GET("/v1", s.handleListUrl())
}
