package routers

import "github.com/trongtb88/rateservice/api/middlewares"

func (s *Server) initializeRoutes() {

	// Health Check Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.)).Methods("GET")
}
