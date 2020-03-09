package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gitlab.com/siwacarn/Golang_API/Farming_API/api/models"
	"gitlab.com/siwacarn/Golang_API/Farming_API/api/responses"
	"gitlab.com/siwacarn/Golang_API/Farming_API/api/utils/formaterror"
)

func (server *Server) CreateSensor(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	sensor := models.Sensor{}
	err = json.Unmarshal(body, &sensor)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	sensorCreated, err := sensor.SaveSensor(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ErrorResponse(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, sensorCreated.Id))
	responses.JsonResponse(w, http.StatusCreated, sensorCreated)
}

func (server *Server) GetSensors(w http.ResponseWriter, r *http.Request) {
	sensor := models.Sensor{}

	sensors, err := sensor.FindAllSensors(server.DB)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	responses.JsonResponse(w, http.StatusOK, sensors)
}

// find date by ISO8601 Format Only (send by `created_at` variable)
func (server *Server) GetSensorByDate(w http.ResponseWriter, r *http.Request) {
	// extract request to body
	body, err := ioutil.ReadAll(r.Body)
	// error handler
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	// create struct of sensor model
	sensor := models.Sensor{}
	// unmarshal body to object
	err = json.Unmarshal(body, &sensor)
	// error handler of marshal
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	// parsing variable to find data
	println(sensor.CreatedAt.String())
	sensorbydate, err := sensor.FindSensorByDate(server.DB, sensor.CreatedAt)
	if err != nil {
		responses.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	responses.JsonResponse(w, http.StatusOK, sensorbydate)
}
