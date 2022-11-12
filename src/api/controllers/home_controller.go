package controllers

import (
	"net/http"

	"github.com/rizal282/api-golang-mux/src/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome...")
}