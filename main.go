package main

import (
	install "FoxxoOS/installation"
	s "FoxxoOS/main_server"
	"fmt"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := exec.Command("bash", "-c", "firefox --new-tab http://localhost:8080")

	if err != nil {
		fmt.Println(err)
	}

	install.Partitioning()

	app := fiber.New(fiber.Config{s
		AppName: "Foxxo OS",
	})

	s.MainServer(app)
}
