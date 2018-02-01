package main

import (
	"fmt"
	"github.com/mitchellh/go-ps"
)

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

	for checkRunning(PUBG) {
		fmt.Println("Game is currently running.")
		fmt.Scanln("Exit game and press enter to continue.")
	}
}
