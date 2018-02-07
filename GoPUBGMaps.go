package main

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"path/filepath"
	"os"
	"runtime"
	"strings"
)

const PUBGExe  = "TslGame.exe"
const GameFolder = "PUBG"
var ChildFolders = [3]string{"_CommonRedist", "Engine", "TslGame"}

const DefaultPath64 = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\PUBG"
const DefaultPath32 = "C:\\Program Files\\Steam\\steamapps\\common\\PUBG"
const RelativeContentPath = "TslGame\\Content\\Paks"

var MapNames = [2]string{"desert", "erangel"}

type Map struct {
	name 	string
	active	bool
	files 	[]string
}

func isRunning(processName string) bool {
	out, _ := ps.Processes()
	for i := 0; i < len(out); i++ {
		if processName == out[i].Executable() {
			return true
		}
	}
	return false
}

func confirmGamePath(rootPath string) bool {
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

func getMapPaths(rootPath string) []string {
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

func parsePathsToMaps(items []string) []Map {
	var maps []Map

	for i := 0; i < len(items); i++ {
		for j := 0; j < len(MapNames); j++ {
			if strings.Contains(items[i], MapNames[j]) {
				for k := 0; k < len(maps); k++ {
					if maps[k].name == MapNames[j] {
						maps[k].files = append(maps[k].files, items[i])
					} else {
						maps = append(maps, Map{name:MapNames[j], active:true, files:[]string{}})
					}
				}
			}
		}
	}
	return maps
}

func getActiveStatus(maps []Map) {
	// TODO: Add code to set map active status from filenames
}

func main() {
	var input string
	var gamePath = ""

	for isRunning(PUBGExe) {
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
		if confirmGamePath(DefaultPath64) {
			gamePath = DefaultPath64
		} else if confirmGamePath(DefaultPath32) {
			gamePath = DefaultPath32
		} else if confirmGamePath(input) {
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
