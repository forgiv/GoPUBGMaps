package main

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const PUBGExe = "TslGame.exe"
const GameFolder = "PUBG"

var ChildFolders = [3]string{"_CommonRedist", "Engine", "TslGame"}

var DefaultPath64 = filepath.Join("C:", "Program Files (x86)", "Steam", "steamapps", "common", "PUBG")
var DefaultPath32 = filepath.Join("C:", "Program Files", "Steam", "steamapps", "common", "PUBG")
var RelativeContentPath = filepath.Join("TslGame", "Content", "Paks")

var MapNames = [2]string{"desert", "erangel"}

type Map struct {
	name   string
	active bool
	files  []string
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

func parsePathsToMaps(items []string) []*Map {
	var maps []*Map

	for i := 0; i < len(MapNames); i++ {
		for j := 0; j < len(items); j++ {
			if len(maps) == 0 {
				maps = append(maps, &Map{name: MapNames[i], active: true, files: []string{items[j]}})
			} else {
				for k := 0; k < len(maps); k++ {
					// TODO: Finish this code
				}
			}
		}
	}
	return maps
}

func updateActiveStatusFromFilenames(maps []*Map) {
	for i := 0; i < len(maps); i++ {
		if filepath.Ext(maps[i].files[0]) == "disabled" {
			maps[i].active = false
		}
	}
}

func setFilenamesFromActiveStatus(maps []*Map) {
	for i := 0; i < len(maps); i++ {
		if maps[i].active {
			for j := 0; j < len(maps[i].files); j++ {
				if filepath.Ext(maps[i].files[j]) == "disabled" {
					maps[i].files[j] = maps[i].files[j][:len(maps[i].files[j])-len("disabled")]
				}
			}
		} else {
			for j := 0; j < len(maps[i].files); j++ {
				if filepath.Ext(maps[i].files[j]) != "disabled" {
					maps[i].files[j] = maps[i].files[j] + ".disabled"
				}
			}
		}
	}
}

func main() {
	var input string
	var gamePath = ""
	var maps []*Map

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

	maps = parsePathsToMaps(getMapPaths(gamePath))
	updateActiveStatusFromFilenames(maps)

	for i := 0; i < len(maps); i++ {
		fmt.Printf("%s: %t\n", maps[i].name, maps[i].active)
	}
}
