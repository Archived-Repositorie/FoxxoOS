package main

import (
	"log"
	"os"
	"flag"

	//install "FoxxoOS/installation"
	s "FoxxoOS/main_server"
	"FoxxoOS/util"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"github.com/gofiber/fiber/v2"
)

func main() {
	util.Clean()
	//install.Installation()

	if back := flag.Bool("backend", false, "Flag for backend"); *back {
		server()
	} else if front := flag.Bool("frontend", false, "Flag for electron"); *front {
		electron()
	} else {
		log.Fatal("Use flag in command! \n  --backend Runs as backend \n  --frontend Runs for frontend (electron)")
	}
}

func electron() {
	elecApp, err := astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName:            "FoxxoOS",
		BaseDirectoryPath:  "foxxoos",
		AppIconDefaultPath: "public/icon/icon.png",
	})
	util.ErrorCheck(err)

	defer elecApp.Close()

	elecApp.HandleSignals()

	err = elecApp.Start()
	util.ErrorCheck(err)

	var window *astilectron.Window
	window, err = elecApp.NewWindow("http://127.0.0.1:8080", &astilectron.WindowOptions{
		Center:         astikit.BoolPtr(true),
		Height:         astikit.IntPtr(1200),
		Width:          astikit.IntPtr(1000),
		Fullscreenable: astikit.BoolPtr(true),
		Fullscreen:     astikit.BoolPtr(false),
	})
	util.ErrorCheck(err)

	err = window.Create()
	util.ErrorCheck(err)

	elecApp.Wait()
}

func server() {
	app := fiber.New(fiber.Config{
		AppName: "Foxxo OS",
	})

	s.MainServer(app)
}
