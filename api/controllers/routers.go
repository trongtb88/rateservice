package controllers

import "github.com/trongtb88/rateservice/api/middlewares"

func (server *Server) InitializeRoutes() {
	// Health Check Route
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.HealthCheck)).Methods("GET")
	server.Router.HandleFunc("/rates/latest", middlewares.SetMiddlewareJSON(server.GetLatestRate)).Methods("GET")
	server.Router.HandleFunc("/rates/{date}", middlewares.SetMiddlewareJSON(server.GetRateSpecificDate)).Methods("GET")
	server.Router.HandleFunc("/rates/analyze", middlewares.SetMiddlewareJSON(server.AnalyzeRates)).Methods("GET")
}
