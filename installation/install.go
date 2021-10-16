package installation

import (
	"FoxxoOS/files"
	"FoxxoOS/util"
	"encoding/json"
	"fmt"
	"os"
)

type Partitions struct {
	Disk string
	Root string
	Swap string
	Boot string
}

func Partitioning() {
	file, err := os.ReadFile(files.FilesJSON[2])

	util.ErrorCheck(err)

	var JSON map[string]map[string]string
	json.Unmarshal(file, &JSON)

	diskInfo := JSON["disk"]

	parts := Partitions{}

	switch diskInfo["type"] {
	case "auto":
		_, err := os.Stat("/sys/firmware/efi")
		fmt.Println(diskInfo)

		rootStart := "0.0"
		if err == nil {
			rootStart = "512M"
		}

		util.Partitioning(
			diskInfo["disk"],
			"mklabel",
			[]string{"gpt"},
			[]string{},
		)
		parts.Disk = diskInfo["disk"]

		partitionRoot := util.Partitioning(
			diskInfo["disk"],
			"mkpart",
			[]string{"primary"},
			[]string{rootStart, "-8G"},
			1,
		)
		parts.Root = partitionRoot

		partitionSwap := util.Partitioning(
			diskInfo["disk"],
			"mkpart",
			[]string{"primary", "linux-swap"},
			[]string{"-8G", "100%"},
			2,
		)
		parts.Swap = partitionSwap

		if err == nil {
			partitionBoot := util.Partitioning(
				diskInfo["disk"],
				"mkpart",
				[]string{"ESP", "fat32"},
				[]string{"-8G", "100%"},
				3,
			)
			parts.Boot = partitionBoot
		}

	case "manual":
		_, err := os.Stat("/sys/firmware/efi")
		if err == nil {
			parts.Boot = diskInfo["boot"]
		}

		parts.Root = diskInfo["root"]
		parts.Swap = diskInfo["swap"]
	}

	fmt.Println(parts)

	util.FormatFS("fs.btrfs", parts.Root)
	util.FormatFS("swap", parts.Swap)

	_, err = os.Stat("/sys/firmware/efi")
	if err == nil {
		util.FormatFS("fs.fat -F 32", parts.Boot)
	}
}
