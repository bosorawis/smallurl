package http

func (s *Server) routes() {
	s.router.POST("/v1", s.handleUrlCreate())
	s.router.POST("/v1/alias", s.handleUrlCreateWithAlias())
	s.router.GET("/v1", s.handleUrlList())
	s.router.GET("/r/:id", s.handleUrlRedirect())
}
