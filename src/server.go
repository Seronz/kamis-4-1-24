package src

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/seronz/api/config"
	"github.com/seronz/api/src/utils/seeder"
	"github.com/urfave/cli"
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
}

func (server *Server) InitializeDB() {
	server.DB = config.DB
}

func (server *Server) dbMigrate() {
	for _, model := range RegistryModels() {
		err := server.DB.Debug().AutoMigrate(model.Models)

		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("migration success...")

	// err := seeder.DBSeed(server.DB)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func (server *Server) initCommands(config AppConfig) {
	server.InitializeDB()
	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeder.DBSeed(server.DB)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "db:delete",
			Action: func(c *cli.Context) error {
				err := seeder.DropAllTable(server.DB)
				if err != nil {
					return err
				}
				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run() {
	config.Connect()
	config.ConnectMongo()
	var server = Server{}
	var appConfig = AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load env")
	}

	appConfig.AppName = os.Getenv("APP_NAME")
	appConfig.AppPort = os.Getenv("APP_PORT")

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.initCommands(appConfig)
	} else {
		server.Initialize()
		server.Run(":" + appConfig.AppPort)
	}

}
