package controllers

import (
	"github.com/trongtb88/rateservice/api"
	_ "github.com/trongtb88/rateservice/api"
	"github.com/trongtb88/rateservice/api/responses"
	"net/http"
)

func (server * Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to rate system")
}
