package middlewares

import (
	"errors"
	"net/http"

	"github.com/siwacarn/Golang_API/Farming_API/api/auth"
	"github.com/siwacarn/Golang_API/Farming_API/api/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ErrorResponse(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
