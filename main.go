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

	// user endpoint
	userRoute := app.Group("/api/user")
	userRoute.Post("/", routes.CreateUser)
	userRoute.Get("/", routes.GetUsers)
	userRoute.Get("/:id", routes.GetUser)
	userRoute.Put("/:id", routes.UpdateUser)
	userRoute.Delete(":id", routes.DeleteUser)

	// product endpoint
	productRoute := app.Group("/api/product")
	productRoute.Get("/", routes.GetProducts)
	productRoute.Post("/", routes.CreateProduct)
	productRoute.Get("/:id", routes.GetProduct)
	productRoute.Put("/:id", routes.PutProduct)
	productRoute.Delete("/:id", routes.DeleteProduct)

	// order endpoint
	orderRoute := app.Group("/api/order")
	orderRoute.Post("/", routes.PostOrder)
	orderRoute.Get("/", routes.GetOrders)
	orderRoute.Get("/:id", routes.GetOrder)

}
