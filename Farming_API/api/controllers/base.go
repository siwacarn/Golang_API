package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/siwacarn/Golang_API/Farming_API/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	if DbDriver == "mysql" {
		DBUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			DbUser,
			DbPassword,
			DbHost,
			DbPort,
			DbName)

		server.DB, err = gorm.Open(DbDriver, DBUri)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", DbDriver)
			log.Fatal("Error: ", err)
		} else {
			fmt.Printf("Database Establish with %s Sucessful! ", DbDriver)
		}
	}

	// Auto migration to database
	server.DB.AutoMigrate(&models.User{})
	// set Router
	server.Router = mux.NewRouter()
	// Initialize Route
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
