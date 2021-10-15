package installation

import (
	"FoxxoOS/files"
	"FoxxoOS/util"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Disk struct {
	Root int
	Swap int
	Boot int
}

func Partitioning() {
	file, err := os.ReadFile(files.FilesJSON[2])

	util.ErrorCheck(err)

	var JSON map[string]map[string]string
	json.Unmarshal(file, &JSON)

	diskInfo := JSON["disk"]

	root, err := strconv.Atoi(diskInfo["root"][len(diskInfo["root"])-1:])
	util.ErrorCheck(err)

	swap, err := strconv.Atoi(diskInfo["swap"][len(diskInfo["swap"])-1:])
	util.ErrorCheck(err)

	boot, err := strconv.Atoi(diskInfo["boot"][len(diskInfo["boot"])-1:])
	util.ErrorCheck(err)

	disk := Disk{
		Root: root,
		Swap: swap,
		Boot: boot,
	}

	switch diskInfo["type"] {
	case "auto":
		_, err := os.Stat("/sys/firmware/efi")
		fmt.Println(disk)

		rootStart := "0.0"
		if err == nil {
			rootStart = "512M"
		}

		partitionRoot := util.Partitioning(
			diskInfo["disk"],
			"mkpart",
			[]string{"primary"},
			[]string{rootStart, "-8G"},
			disk.Root,
		)

		fmt.Println(partitionRoot)

		partitionSwap := util.Partitioning(
			diskInfo["disk"],
			"mkpart",
			[]string{"primary", "linux-swap"},
			[]string{"-8G", "100%"},
			disk.Root,
		)

		if err == nil {
			partitionBoot := util.Partitioning(
				diskInfo["disk"],
				"mkpart",
				[]string{"ESP", "fat32"},
				[]string{"-8G", "100%"},
				disk.Root,
			)

			fmt.Println(partitionBoot)
		}

		fmt.Println(partitionSwap)
	}

}
