package installation

import (
	"FoxxoOS/files"
	"FoxxoOS/util"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)


func Installation() {
	fmt.Println("Startng installation...\n\n")

	fmt.Println("Partitioning...")
	parts := Partitioning()
	fmt.Println("Done!\n\n")

	fmt.Println("Formatting...")
	Formating(parts)
	fmt.Println("Done!\n\n")

	fmt.Println("Mounting...")
	Mounting(parts)
	fmt.Println("Done!\n\n")
}

type Partitions struct {
	Disk string
	Root string
	Swap string
	Boot string
}

func partAuto(parts *Partitions, diskInfo map[string]string) {
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
}

func partManual(parts *Partitions, diskInfo map[string]string) {
	_, err := os.Stat("/sys/firmware/efi")
	if err == nil {
		parts.Boot = diskInfo["boot"]
	}

	parts.Root = diskInfo["root"]
	parts.Swap = diskInfo["swap"]
}

func Formating(parts Partitions) {
	util.FormatFS("fs.btrfs", parts.Root)
	util.FormatFS("swap", parts.Swap)

	_, err := os.Stat("/sys/firmware/efi")
	if err == nil {
		util.FormatFS("fs.fat -F 32", parts.Boot)
	}
}

func Mounting(parts Partitions) {
	util.Mount(parts.Root, "/mnt")

	_, err := os.Stat("/sys/firmware/efi")
	if err == nil {
		command := fmt.Sprintf("mkdir %v", "/mnt/boot")
		cmd := exec.Command("sudo " + command)
		cmd.Run()

		util.Mount(parts.Boot, "/mnt/boot")
	}

	command := fmt.Sprintf("swapon %v", parts.Swap)
	cmd := exec.Command("sudo " + command)
	cmd.Run()
}

func UMounting() {
	_, err := os.Stat("/sys/firmware/efi")
	if err == nil {
		util.UMount("/mnt/boot")
	}

	util.UMount("/mnt")
}

func Partitioning() Partitions {
	file, err := os.ReadFile(files.FilesJSON[2])
	util.ErrorCheck(err)

	var JSON map[string]map[string]string
	json.Unmarshal(file, &JSON)

	diskInfo := JSON["disk"]
	parts := Partitions{}

	switch diskInfo["type"] {
	case "auto":
		partAuto(&parts, diskInfo)
	case "manual":
		partManual(&parts, diskInfo)
	}

	fmt.Println(parts)

	return parts
}

func Config() {
	fileSave, err := os.ReadFile(files.FilesJSON[2])
	util.ErrorCheck(err)

	fileNIX, err := os.ReadFile(files.FilesNIX[0])
	util.ErrorCheck(err)

	var JSON map[string]interface{}
	err = json.Unmarshal(fileSave, &JSON)

	util.ErrorCheck(err)

	start := util.ReplaceFile(fileNIX, "$keyboard", JSON["keyboard"])
	start = util.ReplaceFile(start, "$locales", JSON["lang"])
	start = util.ReplaceFile(start, "$timezone", JSON["timezone"])
	start = util.ReplaceFile(start, "$hostname", JSON["hostname"])

	fmt.Println(string(start))

	fmt.Println(JSON)

	util.SaveFile("nix/test.nix", start)
}