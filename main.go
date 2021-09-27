package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

//im sorry but I gtg Ill be back in aprox 30-40 mins, np im just settuping go file now UwU
//thx :3
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
