package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "poralis"
const collectionName = "games"

var Collection *mongo.Collection

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load environment variables: ", err)
	}
	dbhost := os.Getenv("DB_HOSTNAME")
	dbuser := os.Getenv("DB_USERNAME")
	dbpassword := os.Getenv("DB_PASSWORD")
	dburl := fmt.Sprintf("mongodb://%v:%v@%s:27017", dbuser, dbpassword, dbhost)
	options := options.Client().ApplyURI(dburl)
	client, err := mongo.Connect(context.Background(), options)
	if err != nil {
		log.Fatal("DB connection failed: ", err)
	}
	log.Println("Connected DB")

	Collection = client.Database(dbName).Collection(collectionName)
}
