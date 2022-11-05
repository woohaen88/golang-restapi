package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/woohaen88/database"
	"github.com/woohaen88/routes"
	"log"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}

func HelloWorld(ctx *fiber.Ctx) error {
	msg := fmt.Sprintln("Hello world")
	return ctx.SendString(msg)
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", HelloWorld)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
}
