package main

import (
	s "FoxxoOS/main_server"
	"FoxxoOS/util"
	"fmt"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//Test()

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

func Test() {
	util.Partitioning(
		"/dev/sdc", 
		"mkpart", 
		[]string{"primary",""}, 
		[]string{"-1MiB","100%"},
	)
}