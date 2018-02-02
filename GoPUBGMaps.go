package main

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"path"
	"path/filepath"
	"os"
)

const PUBGExe  = "TslGame.exe"
const GameFolder = "PUBG"

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
	var correctPath = false

	err := filepath.Walk(rootPath, func(filePath string, info os.FileInfo, err error) error {
		if path.Base(filePath) == GameFolder {
			correctPath = true
		}
		return err
	})

	if err != nil {
		panic(err)
	}

	return correctPath
}

func main() {
	var input string

	for checkRunning(PUBGExe) {
		fmt.Println("Game is currently running.")
		fmt.Println("Exit game and press enter to continue.")
		fmt.Scanln(&input)
	}
}
