package main

import (
	//install "FoxxoOS/installation"
	s "FoxxoOS/main_server"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//install.Installation()

	app := fiber.New(fiber.Config{
		AppName: "Foxxo OS",
	})

	s.MainServer(app)
}
