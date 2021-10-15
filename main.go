package main

import (
	s "FoxxoOS/main_server"
	install "FoxxoOS/installation"
	"fmt"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func main() {
	install.Partitioning()

	cmd := exec.Command(
		"bash", 
		"-c", 
		"firefox --new-tab http://127.0.0.1:8080",
	) //--kiosk for full screen installer
	err := cmd.Run()

	if err != nil {
		fmt.Println("Install firefox!")
		fmt.Println(err)
	}

	app := fiber.New(fiber.Config{
		AppName: "Foxxo OS",
	})

	s.MainServer(app)
}