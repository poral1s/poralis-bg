package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/h4y4bus4/polaris/polaris-bg/db"
	"github.com/h4y4bus4/polaris/polaris-bg/handlers"
)

func main() {
	db.Init()

	app := fiber.New()
	handlers.Init(app)

	log.Println(app.Listen(":3000"))
}
