package util

import (
	"log"
	"os"
	"os/exec"
	"FoxxoOS/files"

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

func SaveMain(key string, value string) {
	saveMainRead, err := os.ReadFile(files.Files[2])

	ErrorCheck(err)

	saveMainJSON := string(saveMainRead)
	saveMainJSON,err = sjson.Set(saveMainJSON, key, value)

	ErrorCheck(err)

	err = os.WriteFile(files.Files[2], []byte(saveMainJSON), 0777)

	ErrorCheck(err)
}