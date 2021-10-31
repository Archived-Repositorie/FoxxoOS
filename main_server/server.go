package main_server

import (
	r "FoxxoOS/routes"
	"FoxxoOS/util"

	"github.com/gofiber/fiber/v2"
)

func MainServer(app *fiber.App) {
	app.Get("/post/partition", r.Partitions)
	app.Get("/post/drivers", r.Drivers)
	app.Get("/post/mediagrap", r.MediaGrap)
	app.Get("/post/utils", r.Utils)
	app.Get("/post/gaming", r.Gaming)
	app.Get("/post/office", r.Office)
	app.Get("/post/program", r.Program)
	app.Get("/post/web", r.Web)
	app.Get("/post/de", r.DE)
	app.Get("/post/user", r.User)
	app.Get("/post/timezone", r.Timezone)
	app.Get("/post/lang", r.Lang)
	app.Get("/post/keyboard", r.Keyboard)
	app.Get("/post/install", r.Installation)

	app.Static("/", "./public")
	app.Static("/style", "./style")
	app.Static("/data", "./data")

	err := app.Listen(":8080")

	util.ErrorCheck(err)
}
