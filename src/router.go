package src

import (
	"github.com/seronz/api/src/controllers"
	usercontroller "github.com/seronz/api/src/controllers/user_controller"
	"github.com/seronz/api/src/middleware"
)

func (server *Server) InitializeRouter() {
	server.Router.Use(middleware.LoggerMiddelware)
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
	server.Router.HandleFunc("/register", usercontroller.User).Methods("POST")
	server.Router.HandleFunc("/regenerate_otp", usercontroller.GenerateOTP).Methods("POST")
	server.Router.HandleFunc("/activate", usercontroller.ActivateAccount).Methods("POST")
	server.Router.HandleFunc("/login", usercontroller.Login).Methods("POST")
	server.Router.HandleFunc("/address/create", usercontroller.CreateAddress).Methods("POST")
}
