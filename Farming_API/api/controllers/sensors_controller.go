package controllers

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/siwacarn/Farming_API/api/models"
)

func (server *Server) CreateSensor(c *gin.Context){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ErrorResponse(http.StatusUnprocessableEntity, err)
		return
	}
	sensor := models.Sensor{}
	err = c.BindJSON(&sensor)
	if err != nil {
		responses.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	sensorCreated, err := sensor.SaveSensor(server.DB)

	if err != nil {
		
	}
}