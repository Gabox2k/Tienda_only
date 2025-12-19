package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func conexion() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cliente, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))

	if err != nil {
		log.Fatal(err)
	}

	DB = cliente.Database("pinguinos")
	log.Println("mongodb conectado")

}
