package server


func (s *Server) routes() {
	s.router.MethodFunc("PUT","/v1", s.handlePut())
	s.router.MethodFunc("GET", "/v1", s.handleList())
}
