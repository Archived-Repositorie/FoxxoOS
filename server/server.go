package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
) //"github.com/tidwall/sjson"

func Index(c *fiber.Ctx) error {
	return c.SendFile("./public/index.html")
}

func Lang(c *fiber.Ctx) error {
	files := []string{"data/languages.json"}

	langRead, _ := os.ReadFile(files[0])
	langJSON := string(langRead)
	lang := c.Query("lang")

	value := gjson.Get(langJSON, lang)
	fmt.Println(value)

	return c.SendString(value.String())
}
