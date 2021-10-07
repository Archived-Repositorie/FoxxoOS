package server

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
) //

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

var files = [...]string{
	"data/languages.json", 
	"data/keyboard.json", 
	"data/save.json",
	"data/timezones.json",
}

func errorCheck(er error) {
	if er != nil {
		log.Fatalln(er)

		cmd := exec.Command("bash", "-c", "killall firefox")
		cmd.Run()
	}
}

func Save(c *fiber.Ctx) error {
	saveRead, err := os.ReadFile(files[2])
	
	errorCheck(err)

	saveJSON := string(saveRead)

	if c.Query("done") != "ok" {
		return c.SendString("not ok")
	}

	return c.SendString(saveJSON)
}

func SaveMain(key string, value string) {
	saveMainRead, err := os.ReadFile(files[2])

	errorCheck(err)

	saveMainJSON := string(saveMainRead)
	saveMainJSON,err = sjson.Set(saveMainJSON, key, value)

	errorCheck(err)

	err = os.WriteFile(files[2], []byte(saveMainJSON), 0777)

	errorCheck(err)
}

func Lang(c *fiber.Ctx) error {
	langRead, err := os.ReadFile(files[0])

	errorCheck(err)

	langJSON := string(langRead)
	lang := c.Query("lang")

	value := gjson.Get(langJSON, lang)

	SaveMain("lang", value.String())

	return c.SendString(value.String())
}

func Keyboard(c *fiber.Ctx) error {
	keyRead, err := os.ReadFile(files[1])

	errorCheck(err)

	keyJSON := string(keyRead)
	key := c.Query("keyboard")

	value := gjson.Get(keyJSON, key)

	SaveMain("keyboard", value.String())

	return c.SendString(value.String())
}

type Time struct {
	Timezone []string
}

func Timezone(c *fiber.Ctx) error {
	timeRead, err := os.ReadFile(files[3])

	errorCheck(err)

	var times Time

	err = json.Unmarshal(timeRead, &times)

	errorCheck(err)

	time := c.Query("time")

	if !stringInSlice(time, times.Timezone) {
		return c.SendString("no ok")
	}

	SaveMain("timezone", time)

	return c.SendString(time)
}

func MainServer(app *fiber.App) {
	app.Post("/post/timezone", Timezone)
	app.Post("/post/lang", Lang)
	app.Post("/post/keyboard", Keyboard)
	app.Post("/post/save", Save)

	app.Static("/", "./public")
	app.Static("/style", "./style")

	err := app.Listen(":8080")

	errorCheck(err)
}
