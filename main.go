package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sumits31/fiberdemo/user"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome to fleek")
}

func Routers(app *fiber.App) {
	app.Get("/users", user.Getusers)
	app.Post("/users", user.Addusers)
	app.Get("/users/:id", user.Getuser)
	app.Put("/users/:id", user.Updateusers)
	app.Delete("/users/:id", user.Deleteusers)
}

func main() {
	user.InitialMigrations()

	app := fiber.New()

	Routers(app) // Set up routes before starting the server
	app.Get("/", hello)

	app.Listen(":7000") // Start the server directly using Listen
}
 