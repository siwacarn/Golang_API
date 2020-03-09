package controllers

import "gitlab.com/siwacarn/Golang_API/Farming_API/middlewares"

func (s *Server) initializeRoutes() {
	// Home route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	// Users
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	// Sensor
	s.Router.HandleFunc("/sensors", middlewares.SetMiddlewareJSON(s.CreateSensor)).Methods("POST")
	s.Router.HandleFunc("/sensors", middlewares.SetMiddlewareJSON(s.GetSensors)).Methods("GET")
	s.Router.HandleFunc("/sensors/date", middlewares.SetMiddlewareJSON(s.GetSensorByDate)).Methods("POST")
}
