package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbHost = "10.0.2.2"
const mongoUri = "mongodb://root:example@" + dbHost + ":27017"
const dbName = "poralis"

const collectionName = "games"

var Collection *mongo.Collection

func Init() {
	options := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.Background(), options)
	if err != nil {
		log.Fatal("DB connection failed: ", err)
	}
	log.Println("Connected DB")

	Collection = client.Database(dbName).Collection(collectionName)
}
