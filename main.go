package main

import (
	"log"

	"github.com/Hayabusa58/polar1s/polaris-bg/db"
	"github.com/Hayabusa58/polar1s/polaris-bg/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db.Init()

	app := fiber.New()
	app.Use(cors.New())

	handlers.Init(app)

	log.Println(app.Listen(":3000"))
}
