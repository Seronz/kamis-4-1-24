package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mongo *mongo.Client
var globalCtx = context.Background()

func ConnectMongo() {
	client, err := mongo.Connect(globalCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancle := context.WithTimeout(globalCtx, 10*time.Second)
	defer cancle()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	Mongo = client
	fmt.Println("Connection Mongo Success...")
}
