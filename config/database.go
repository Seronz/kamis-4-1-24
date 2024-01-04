package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigConnection struct {
	host     string
	user     string
	password string
	dbname   string
	port     string
}

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to read env")
	}

	var config = ConfigConnection{}

	config.host = os.Getenv("HOST")
	config.user = os.Getenv("USER")
	config.password = os.Getenv("DATABASE_PASSWORD")
	config.dbname = os.Getenv("DATABASE")
	config.port = os.Getenv("DATABASE_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.host,
		config.user,
		config.password,
		config.dbname,
		config.port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("failed to open database")
		panic(err)
	}
	log.Println("Database Connect")
	DB = db
}
