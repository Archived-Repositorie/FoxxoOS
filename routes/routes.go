package routes

import (
	"encoding/json"
	"fmt"
	"os"

	"FoxxoOS/files"
	"FoxxoOS/util"

	"github.com/gofiber/fiber/v2"
	"github.com/tidwall/gjson"
)

func User(c *fiber.Ctx) error {
	user := [3]string{
		c.Query("name"),
		c.Query("password"),
		c.Query("hostname"),
	}

	userJSON := fmt.Sprintf("[%v,%v,%v]", user[0], user[1], user[2])

	util.SaveMain("user.name", user[0])
	util.SaveMain("user.password", user[1])
	util.SaveMain("hostname", user[2])

	return c.SendString(userJSON)
}

func Save(c *fiber.Ctx) error {
	saveRead, err := os.ReadFile(files.Files[2])
	
	util.ErrorCheck(err)

	saveJSON := string(saveRead)

	if c.Query("done") != "ok" {
		return c.SendString("not ok")
	}

	return c.SendString(saveJSON)
}

func Lang(c *fiber.Ctx) error {
	langRead, err := os.ReadFile(files.Files[0])

	util.ErrorCheck(err)

	langJSON := string(langRead)
	lang := c.Query("lang")

	value := gjson.Get(langJSON, lang)

	util.SaveMain("lang", value.String())

	return c.SendString(value.String())
}

func Keyboard(c *fiber.Ctx) error {
	keyRead, err := os.ReadFile(files.Files[1])

	util.ErrorCheck(err)

	keyJSON := string(keyRead)
	key := c.Query("keyboard")

	value := gjson.Get(keyJSON, key)

	util.SaveMain("keyboard", value.String())

	return c.SendString(value.String())
}

type Time struct {
	Timezone []string
}

func Timezone(c *fiber.Ctx) error {
	timeRead, err := os.ReadFile(files.Files[3])

	util.ErrorCheck(err)

	var times Time

	err = json.Unmarshal(timeRead, &times)

	util.ErrorCheck(err)

	time := c.Query("time")

	if !util.StringInSlice(time, times.Timezone) {
		return c.SendString("no ok")
	}

	util.SaveMain("timezone", time)

	return c.SendString(time)
}