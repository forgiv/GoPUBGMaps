package main

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
)

// TslGame.exe

func checkRunning(processName string) bool {
	out, _ := ps.Processes()
	for i := 0; i < len(out); i++ {
		if processName == out[i].Executable() {
			return true
		}
	}
	return false
}

func main() {
	var PUBG = "TslGame.exe"

	if checkRunning(PUBG) {
		fmt.Println("Game is currently running. Exiting...")
		os.Exit(0)
	}
}
