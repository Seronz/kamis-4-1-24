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

	if config.DB == nil {
		log.Println("error ini nil")
	}

	server.DB = config.DB
	for _, model := range RegistryModels() {
		err := server.DB.Debug().AutoMigrate(model.Models)

		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("migration success...")

}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run() {
	config.Connect()
	var server = Server{}
	var appConfig = AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load env")
	}

	appConfig.AppName = os.Getenv("APP_NAME")
	appConfig.AppPort = os.Getenv("APP_PORT")

	server.Initialize()
	server.Run(":" + appConfig.AppPort)
}
