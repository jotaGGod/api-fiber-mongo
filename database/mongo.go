package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client //Variável global mongo db

func ConfigureMongo() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, _ := mongo.Connect(context.TODO(), clientOptions)
	MongoDB = client

	indexModel := mongo.IndexModel{
		Keys: bson.M{"id": 1}, // Escolhemos a chave "id" para o índice
	}

	database := MongoDB.Database("Banco-Atlas").Collection("userAccount")
	_, err := database.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("Connected to MongoDB!")
}

func GetCollection(collectionName string) *mongo.Collection {
	return MongoDB.Database("Banco-Atlas").Collection(collectionName)
}
