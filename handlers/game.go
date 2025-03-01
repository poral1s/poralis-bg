package handlers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/h4y4bus4/polaris/polaris-bg/db"
	"github.com/h4y4bus4/polaris/polaris-bg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getGames(c *fiber.Ctx) error {
	var games []models.Game

	cursor, err := db.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed getting data"})
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var game models.Game
		err := cursor.Decode(&game)
		if err != nil {
			log.Println("Failed decoding data: ", err)
			return c.Status(500).JSON(fiber.Map{"error": "Failed decoding data"})
		}
		games = append(games, game)
	}

	return c.JSON(games)
}

func getGame(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id format: ", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid id format"})
	}

	var game models.Game
	err = db.Collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&game)
	if err != nil {
		log.Println("Could not find specified game: ", err)
		return c.Status(404).JSON(fiber.Map{"error": "Could not find specified game"})
	}

	return c.JSON(game)
}

func createGame(c *fiber.Ctx) error {
	game := new(models.Game)
	err := c.BodyParser(game)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})

	}

	game.ID = primitive.NewObjectID()
	_, err = db.Collection.InsertOne(context.Background(), game)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed saving data"})
	}

	return c.Status(200).JSON(game)
}
