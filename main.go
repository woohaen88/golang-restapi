package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", HelloWorld)

	log.Fatal(app.Listen(":8000"))
}

func HelloWorld(ctx *fiber.Ctx) error {
	msg := fmt.Sprintln("Hello world")
	return ctx.SendString(msg)
}
