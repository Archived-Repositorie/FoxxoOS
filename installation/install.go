package installation

import (
	"FoxxoOS/files"
	"FoxxoOS/util"
	"encoding/json"
	"fmt"
	"os"
)

func Partitioning() {
	file, err := os.ReadFile(files.FilesJSON[2])

	util.ErrorCheck(err)

	var JSON map[string]interface{}
	json.Unmarshal(file, &JSON)

	disk := JSON["disk"]

	fmt.Println(disk)

	// partition := util.Partitioning(
	// 	"/dev/sdc", 
	// 	"mkpart", 
	// 	[]string{"primary"}, 
	// 	[]string{"-1M","100%"},
	// 	3,
	// )

	// fmt.Println(partition)
}