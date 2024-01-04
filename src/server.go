package src

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/seronz/api/config"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppPort string
}

func (server *Server) Initialize() {
	fmt.Println("welcome")
	server.Router = mux.NewRouter()
	server.InitializeRouter()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	config.Connect()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load env")
	}

	appConfig.AppName = os.Getenv("APP_NAME")
	appConfig.AppPort = os.Getenv("APP_PORT")

	server.Initialize()
	server.Run(":" + appConfig.AppPort)
}
