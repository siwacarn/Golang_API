package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/siwacarn/Golang_API/Farming_API/api/auth"
	"github.com/siwacarn/Golang_API/Farming_API/api/models"
	"github.com/siwacarn/Golang_API/Farming_API/api/responses"
	"github.com/siwacarn/Golang_API/Farming_API/api/utils/formaterror"

	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	// make structs of users object
	user := models.User{}
	// unpack request to object of struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Preparing data
	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := server.SignIn(user.Username, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	responses.JsonResponse(w, http.StatusOK, token)
}

func (server *Server) SignIn(username, password string) (string, error) {
	// make variable err
	var err error
	// make object of User from struct
	user := models.User{}

	err = server.DB.Model(models.User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.Id)
}
