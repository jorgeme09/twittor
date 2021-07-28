package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://jorge:jorge007@cluster0.dep7p.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa a la DB")

	return client
}

func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}
	return 1
}
