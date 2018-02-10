package libs

import (
	"path/filepath"
	"os"
	"fmt"
	"strings"
)

var MapNames = [2]string{"desert", "erangel"}

type Map struct {
	Name   string
	Active bool
	Files  []string
}

type Game struct {
	Maps	[]*Map
}

func (game *Game) ParsePathsToMaps(items []string) {
	for i := 0; i < len(MapNames); i++ {
		for j := 0; j < len(items); j++ {
			if strings.Contains(items[j], MapNames[i]) {
				if len(game.Maps) == 0 {
					game.Maps = append(game.Maps, &Map{Name:MapNames[i], Active:true, Files:[]string{items[j]}})
				} else {
					var found = false
					for k := 0; k < len(game.Maps); k++ {
						if game.Maps[k].Name == MapNames[i] {
							game.Maps[k].Files = append(game.Maps[k].Files, items[j])
							found = true
						}
					}
					if !found {
						game.Maps = append(game.Maps, &Map{Name:MapNames[i], Active:true, Files:[]string{items[j]}})
					}
				}
			}
		}
	}
}

func (game *Game) UpdateActiveStatusFromFilenames() {
	for i := 0; i < len(game.Maps); i++ {
		if filepath.Ext(game.Maps[i].Files[0]) == ".disabled" {
			game.Maps[i].Active = false
		}
	}
}

func (game *Game) SetFilenamesFromActiveStatus() {
	for i := 0; i < len(game.Maps); i++ {
		if game.Maps[i].Active {
			for j := 0; j < len(game.Maps[i].Files); j++ {
				if filepath.Ext(game.Maps[i].Files[j]) == ".disabled" {
					newName := game.Maps[i].Files[j][:len(game.Maps[i].Files[j])-len(".disabled")]
					err := os.Rename(game.Maps[i].Files[j], newName)
					if err != nil {
						fmt.Printf("Error renaming file. Send this error to maintainer for fixing.\n%s\n", err)
					} else {
						game.Maps[i].Files[j] = newName
					}
				}
			}
		} else {
			for j := 0; j < len(game.Maps[i].Files); j++ {
				if filepath.Ext(game.Maps[i].Files[j]) != ".disabled" {
					newName := game.Maps[i].Files[j] + ".disabled"
					err := os.Rename(game.Maps[i].Files[j], newName)
					if err != nil {
						fmt.Printf("Error renaming file. Send this error to maintainer for fixing.\n%s\n", err)
					} else {
						game.Maps[i].Files[j] = newName
					}
				}
			}
		}
	}
}