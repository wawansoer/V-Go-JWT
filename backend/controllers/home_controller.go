package controllers

import (
	"net/http"

	"github.com/wawansoer/V-Go-JWT/backend/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
