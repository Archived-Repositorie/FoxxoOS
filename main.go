package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func index(c *fiber.Ctx) error {
	return c.SendFile("./public/index.html")
}

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Test Fiber",
	})

	app.Static("/style", "./style")

	app.Get("/", index)

	err := app.Listen(":8080")
	fmt.Println(err)
}
