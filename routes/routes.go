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

	userJSON := fmt.Sprintf("%v", user)

	util.SetOnceSave("user.name", user[0])
	util.SetOnceSave("user.password", user[1])
	util.SetOnceSave("hostname", user[2])

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

	util.SetOnceSave("lang", value.String())

	return c.SendString(value.String())
}

func Keyboard(c *fiber.Ctx) error {
	keyRead, err := os.ReadFile(files.Files[1])

	util.ErrorCheck(err)

	keyJSON := string(keyRead)
	key := c.Query("keyboard")

	value := gjson.Get(keyJSON, key)

	util.SetOnceSave("keyboard", value.String())

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

	util.SetOnceSave("timezone", time)

	return c.SendString(time)
}

func DE(c *fiber.Ctx) error {
	DERead, err := os.ReadFile(files.Files[4])

	util.ErrorCheck(err)

	DEJSON := string(DERead)
	DE := c.Query("desktop")

	value := gjson.Get(DEJSON, DE)

	util.SetOnceSave("desktop", value.String())

	return c.SendString(value.String())
}

func Web(c *fiber.Ctx) error {
	webRead, err := os.ReadFile(files.Files[5])

	util.ErrorCheck(err)

	webJSON := string(webRead)
	web := c.Query("web")

	value := gjson.Get(webJSON, web)

	util.SetOnceSave("webbrowser", value.String())

	return c.SendString(value.String())
}

func Program(c *fiber.Ctx) error {
	webRead, err := os.ReadFile(files.Files[6])

	util.ErrorCheck(err)

	webJSON := string(webRead)
	web := c.Query("program")

	value := gjson.Get(webJSON, web)

	util.SetOnceSave("programming", value.String())

	return c.SendString(value.String())
}

func Office(c *fiber.Ctx) error {
	officeRead, err := os.ReadFile(files.Files[7])

	util.ErrorCheck(err)

	officeJSON := string(officeRead)

	var officeMap map[string]string
	json.Unmarshal(officeRead, &officeMap)

	office := []string{}
	for key,_ := range officeMap {
		office = append(office, c.Query(key))
	}

	fmt.Println(office)

	officeList := []string{}
	for i := 0; i < len(office); i++ {
		if office[i] != "" {
			officeList = append(officeList,gjson.Get(officeJSON, office[i]).String())
		}
	}

	util.SetMultiSave("office", officeList)
	
	return c.SendString(fmt.Sprintf("%v", officeList))
}