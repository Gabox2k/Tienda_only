package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// Conexion con MongoDB
func Conexion() {
	//Limite de 10s
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//Libera recursos
	defer cancel()

	//Conecta con MongoDB
	cliente, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))

	if err != nil {
		log.Fatal("Error al conectar ", err)
	}

	//Crea la base de datos
	DB = cliente.Database("pinguinos")
	log.Println("mongodb conectado")

	//Creacion de una lista y guarda en uns BSON
	productos := []interface{}{
		bson.D{{"nombre", "teclado"}, {"precio", 5}, {"imagen", "/imagen/Teclado.jpg"}},
		bson.D{{"nombre", "teclado"}, {"precio", 32}, {"imagen", "/imagen/Teclado2.jpg"}},
		bson.D{{"nombre", "teclado"}, {"precio", 60}, {"imagen", "/imagen/Teclado3.jpg"}},
		bson.D{{"nombre", "Combo teclado + Mouse"}, {"precio", 90}, {"imagen", "/imagen/Combo.jpg"}},
		bson.D{{"nombre", "Mouse"}, {"precio", 10}, {"imagen", "/imagen/Mouse.jpg"}},
		bson.D{{"nombre", "Mouse"}, {"precio", 46}, {"imagen", "/imagen/Mouse2.jpg"}},
		bson.D{{"nombre", "Mouse"}, {"precio", 50}, {"imagen", "/imagen/Mouse3.jpg"}},
		bson.D{{"nombre", "Mouse"}, {"precio", 35}, {"imagen", "/imagen/Mouse4.jpg"}},
		bson.D{{"nombre", "Mouse Inalambrico"}, {"precio", 75}, {"imagen", "/imagen/Mouse5.jpg"}},
		bson.D{{"nombre", "PC"}, {"precio", 500}, {"imagen", "/imagen/PC.jpg"}},
		bson.D{{"nombre", "PC"}, {"precio", 700}, {"imagen", "/imagen/PC2.jpg"}},
		bson.D{{"nombre", "PC"}, {"precio", 1000}, {"imagen", "/imagen/PC3.jpg"}},
	}

	//Cuenta cuanto hay en productos
	count, _ := DB.Collection("productos").CountDocuments(ctx, bson.D{})
	if count == 0 {
		DB.Collection("productos").InsertMany(ctx, productos)
		log.Println("productos creados")
	}

}
