package main

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"GoPUBGMaps/libs"
)

const PUBGExe = "TslGame.exe"
const GameFolder = "PUBG"

var ChildFolders = [3]string{"_CommonRedist", "Engine", "TslGame"}

var DefaultPath64 = filepath.Join("C:", "Program Files (x86)", "Steam", "steamapps", "common", "PUBG")
var DefaultPath32 = filepath.Join("C:", "Program Files", "Steam", "steamapps", "common", "PUBG")
var RelativeContentPath = filepath.Join("TslGame", "Content", "Paks")

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

func main() {
	var input string
	var gamePath = ""
	var game = &libs.Game{}

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



	game.MapsFromPaths(getMapPaths(gamePath))
	game.UpdateMaps()

	for {
		i := 1
		for key, value := range game.GetMaps() {
			fmt.Printf("%d. %s: %t\n", i, key, value)
			i++
		}

		fmt.Println("Enter the number of the map to enable/disable it. Or enter q to quit.")
		fmt.Scanln(&input)

		if input == "q" || input == "Q" {
			break
		}

		mapNum, err := strconv.ParseInt(input, 10, 16)

		if err != nil {
			fmt.Println("Input was invalid. Try again.")
			continue
		}

		game.ToggleActive(int(mapNum-1))
	}
}
