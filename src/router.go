package src

import "github.com/seronz/api/src/controllers"

func (server *Server) InitializeRouter() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
