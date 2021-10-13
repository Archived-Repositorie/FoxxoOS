package util

import (
	"FoxxoOS/files"
	"log"
	"os"
	"os/exec"

	"github.com/tidwall/sjson"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func ErrorCheck(er error) {
	if er != nil {
		log.Fatalln(er)

		cmd := exec.Command("bash", "-c", "killall firefox")
		cmd.Run()
	}
}

func SetMultiSave(key string, value []string) {
	SetMultiSave, err := os.ReadFile(files.FilesJSON[2])

	ErrorCheck(err)

	SetMultiSaveJSON := string(SetMultiSave)
	SetMultiSaveJSON, err = sjson.Set(SetMultiSaveJSON, key, value)

	ErrorCheck(err)

	err = os.WriteFile(files.FilesJSON[2], []byte(SetMultiSaveJSON), 0777)

	ErrorCheck(err)
}

func SetOnceSave(key string, value string) {
	SetOnceSave, err := os.ReadFile(files.FilesJSON[2])

	ErrorCheck(err)

	SetSaveJSON := string(SetOnceSave)
	SetSaveJSON, err = sjson.Set(SetSaveJSON, key, value)

	ErrorCheck(err)

	err = os.WriteFile(files.FilesJSON[2], []byte(SetSaveJSON), 0777)

	ErrorCheck(err)
}
