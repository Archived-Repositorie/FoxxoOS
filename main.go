package main

import (
	"log"
	"sync"
	"os"

	//install "FoxxoOS/installation"
	s "FoxxoOS/main_server"
	"FoxxoOS/util"

	"github.com/gofiber/fiber/v2"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func main() {
	util.Clean()
	//install.Installation()
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go server(wg)

	wg.Add(1)
	go electron(wg)

	wg.Wait()
}

func electron(wg *sync.WaitGroup) {
	elecApp, err := astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName:           "FoxxoOS",
		BaseDirectoryPath: "foxxoos",
		AppIconDefaultPath: "public/icon/icon.png",
	})
	util.ErrorCheck(err)

	defer elecApp.Close()

	elecApp.HandleSignals()

	err = elecApp.Start()
	util.ErrorCheck(err)

	var window *astilectron.Window
	window, err = elecApp.NewWindow("http://127.0.0.1:8080", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(1200),
		Width: astikit.IntPtr(1000),
		Fullscreenable: astikit.BoolPtr(true),
		Fullscreen: astikit.BoolPtr(false),
	})
	util.ErrorCheck(err)

	err = window.Create()
	util.ErrorCheck(err)

	elecApp.Wait()

	wg.Done()
}

func server(wg *sync.WaitGroup) {
	app := fiber.New(fiber.Config{
		AppName: "Foxxo OS",
	})

	s.MainServer(app)

	wg.Done()
}