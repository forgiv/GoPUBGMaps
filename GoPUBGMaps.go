package main

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"path/filepath"
	"os"
	"runtime"
)

const PUBGExe  = "TslGame.exe"
const GameFolder = "PUBG"
var ChildFolders = [3]string{"_CommonRedist", "Engine", "TslGame"}
const DefaultPath64 = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\PUBG"
const DefaultPath32 = "C:\\Program Files\\Steam\\steamapps\\common\\PUBG"
const RelativeContentPath = "TslGame\\Content\\Paks"

func checkRunning(processName string) bool {
	out, _ := ps.Processes()
	for i := 0; i < len(out); i++ {
		if processName == out[i].Executable() {
			return true
		}
	}
	return false
}

func confirmPath(rootPath string) bool {
	var foundChildren = 0

	if filepath.Base(rootPath) == GameFolder {
		err := filepath.Walk(rootPath, func(filePath string, info os.FileInfo, err error) error {
			for i := 0; i < len(ChildFolders); i++ {
				if ChildFolders[i] == filepath.Base(filePath) {
					foundChildren++
				}
			}

			return err
		})

		if err != nil {
			return false
		}
	}

	if foundChildren == len(ChildFolders) {
		return true
	}
	return false
}

func getContent(rootPath string) []string {
	contentPath := filepath.Join(rootPath, RelativeContentPath)
	var items []string
	err := filepath.Walk(contentPath, func(path string, info os.FileInfo, err error) error {
		items = append(items, path)
		return err
	})
	if err != nil {
		panic(err)
	}
	return items
}

func interpretContent(items []string) {
	// TODO: Add code to read filenames and determine which maps are activated/deactived
}

func main() {
	var input string
	var gamePath = ""

	for checkRunning(PUBGExe) {
		fmt.Println("Game is currently running.")
		fmt.Println("Exit game and press enter to continue.")
		fmt.Scanln(&input)
	}

	fmt.Printf("Enter the absolute path of %s folder.\n", GameFolder)
	fmt.Print("Leave blank to use default path: ")
	if runtime.GOARCH == "amd64" {
		fmt.Println(DefaultPath64)
	} else if runtime.GOARCH == "386" {
		fmt.Println(DefaultPath32)
	} else {
		fmt.Println("OS architecture could not be found. No default path set.")
	}

	for {
		fmt.Scanln(&input)
		if confirmPath(DefaultPath64) {
			gamePath = DefaultPath64
		} else if confirmPath(DefaultPath32) {
			gamePath = DefaultPath32
		} else if confirmPath(input) {
			gamePath = input
		}
		if gamePath != "" {
			break
		} else {
			fmt.Println("Folder could not be found. Is game installed? Is correct folder chosen?")
			fmt.Printf("Provide absolute path to %s folder.\n", GameFolder)
			continue
		}
	}
}
