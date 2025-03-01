package handlers

import "github.com/gofiber/fiber/v2"

func Init(app *fiber.App) {
	app.Post("/games", createGame)
	app.Get("/games", getGames)
	app.Get("/games/:id", getGame)
}
