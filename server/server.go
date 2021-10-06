package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
) //"github.com/tidwall/sjson"

func Lang(c *fiber.Ctx) error {
	files := []string{"data/languages.json"}

	langRead, _ := os.ReadFile(files[0])
	langJSON := string(langRead)
	lang := c.Query("lang")

	value := gjson.Get(langJSON, lang)
	fmt.Println(value)

	return c.SendString(value.String())
}

func Keyboard(c *fiber.Ctx) error {
	files := []string{"data/keyboard.json"}

	keyRead, _ := os.ReadFile(files[0])
	keyJSON := string(keyRead)
	key := c.Query("keyboard")

	value := gjson.Get(keyJSON, key)
	fmt.Println(value)

	return c.SendString(value.String())
}

func MainServer(app *fiber.App) {
	app.Static("/style", "./style")

	app.Post("/post/lang", Lang)
	app.Post("/post/keyboard", Keyboard)
	app.Static("/", "./public")

	err := app.Listen(":8080")
	fmt.Println(err)
}
