package src

import (
	"github.com/seronz/api/src/controllers"
	usercontroller "github.com/seronz/api/src/controllers/user_controller"
)

func (server *Server) InitializeRouter() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
	server.Router.HandleFunc("/register", usercontroller.User).Methods("POST")
}
