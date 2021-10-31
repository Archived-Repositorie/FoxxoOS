package main

import (
	//install "FoxxoOS/installation"
	s "FoxxoOS/main_server"
	"FoxxoOS/util"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//install.Installation()
	util.Clean()
	app := fiber.New(fiber.Config{
		AppName: "Foxxo OS",
	})

	s.MainServer(app)
}
