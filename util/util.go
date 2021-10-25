package util

import (
	"FoxxoOS/files"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/tidwall/sjson"
)

func StringInSlice(a string, list interface{}) bool {
	switch list := list.(type) {
	case []interface{}:
		for _, b := range list {
			if b == a {
				return true
			}
		}
	case []string:
		for _, b := range list {
			if b == a {
				return true
			}
		}
	}
	return false
}

func GetString(list interface{}, key string) string {
	switch list := list.(type) {
	case map[string]interface{}:
		return fmt.Sprintf("%v", list[key])
	}
	return "user"
}

func ErrorCheck(er error) {
	if er != nil {
		log.Fatalln(er)

		cmd := exec.Command("bash", "-c", "killall firefox")
		cmd.Run()
	}
}

func Stringing(array interface{}, replace string) string {
	var str string
	switch array := array.(type) {
	case []interface{}:
		for _, a := range array {
			str += fmt.Sprintf("%v%v", a, replace)
		}
	}
	return str
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

func Partitioning(disk string, option string, types []string, start_end []string, partition ...int) string {
	command := fmt.Sprintf(
		"parted -s %v -- %v %v %v",
		disk,
		option,
		strings.Join(types, " "),
		strings.Join(start_end, " "),
	)

	cmd := exec.Command("bash", "-c", "sudo "+command)

	err := cmd.Run()

	ErrorCheck(err)

	return fmt.Sprintf("%v%v", disk, partition)
}

func FormatFS(fs string, partition string) string {
	command := fmt.Sprintf(
		"mk%v -f %v",
		fs,
		partition,
	)

	cmd := exec.Command("bash", "-c", "sudo "+command)

	err := cmd.Run()

	ErrorCheck(err)

	return fmt.Sprintf("%v %v", partition, fs)
}

func Mount(partition string, folder string) string {
	command := fmt.Sprintf(
		"mount %v %v",
		partition,
		folder,
	)

	cmd := exec.Command("bash", "-c", "sudo "+command)

	err := cmd.Run()

	ErrorCheck(err)

	return fmt.Sprintf("%v %v", partition, folder)
}

func UMount(folder string) string {
	command := fmt.Sprintf(
		"umount %v",
		folder,
	)

	cmd := exec.Command("bash", "-c", "sudo "+command)

	err := cmd.Run()

	ErrorCheck(err)

	return folder
}

func ReplaceFile(file *[]byte, key string, value interface{}) {
	stringFile := string(*file)

	replaceString := strings.ReplaceAll(stringFile, key, fmt.Sprintf("%v", value))

	*file = []byte(replaceString)
}

func SaveFile(file string, value []byte) {
	err := os.WriteFile(file, value, 0777)

	ErrorCheck(err)
}

func SudoExec(command string, add ...interface{}) {
	command = fmt.Sprintf(command, add...)

	cmd := exec.Command("bash", "-c", "sudo "+command)

	fmt.Println(command)

	err := cmd.Run()

	ErrorCheck(err)
}
