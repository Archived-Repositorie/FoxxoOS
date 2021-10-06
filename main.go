package main

import (
	"FoxxoOS/server"
	"fmt"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cmd := exec.Command("bash", "-c", "firefox --new-tab http://127.0.0.1:8080") //--kiosk for full screen installer

	err := cmd.Run()

	if err != nil {
		fmt.Println("Install firefox!")
		fmt.Println(err)
	}

	app := fiber.New(fiber.Config{
		AppName: "Foxxo OS",
	})

	app.Static("/style", "./style")

	app.Post("/post/lang", server.Lang)
	app.Get("/", server.Index)

	err = app.Listen(":8080")
	fmt.Println(err)
}
