package controllers

import (
	"net/http"

	"github.com/siwacarn/Golang_API/Farming_API/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JsonResponse(w, http.StatusOK, "Farming API v1.0 - release version")
}
