package main_server

import (
	r "FoxxoOS/routes"
	"FoxxoOS/util"

	"github.com/gofiber/fiber/v2"
)


func MainServer(app *fiber.App) {
	app.Post("/post/utils", r.Utils)
	app.Post("/post/gaming", r.Gaming)
	app.Post("/post/office", r.Office)
	app.Post("/post/program", r.Program)
	app.Post("/post/web", r.Web)
	app.Post("/post/de", r.DE)
	app.Post("/post/user", r.User)
	app.Post("/post/timezone", r.Timezone)
	app.Post("/post/lang", r.Lang)
	app.Post("/post/keyboard", r.Keyboard)
	app.Post("/post/save", r.Save)

	app.Static("/", "./public")
	app.Static("/style", "./style")

	err := app.Listen(":8080")

	util.ErrorCheck(err)
}
