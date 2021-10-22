package main

import (
	s "FoxxoOS/main_server"
	// install "FoxxoOS/installation"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// install.Config()

	app := fiber.New(fiber.Config{
		AppName: "Foxxo OS",
	})

	s.MainServer(app)
}
