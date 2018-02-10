package libs

import (
	"path/filepath"
	"os"
	"fmt"
	"strings"
)

var mapNames = [2]string{"desert", "erangel"}

type Map struct {
	Name   string
	Active bool
	Files  []string
}

func (m *Map) activeFromFiles() {
	m.Active = filepath.Ext(m.Files[0]) == ".disabled"
}

func (m *Map) enableFiles() {
	for i := 0; i < len(m.Files); i++ {
		newName := m.Files[i][:len(m.Files[i])-len(".disabled")]
		err := os.Rename(m.Files[i], newName)
		if err != nil {
			fmt.Printf("Error renaming file: %s\n", err)
		} else {
			m.Files[i] = newName
		}
	}
}

func (m *Map) disableFiles() {
	for i := 0; i < len(m.Files); i++ {
		newName := m.Files[i] + ".disabled"
		err := os.Rename(m.Files[i], newName)
		if err != nil {
			fmt.Printf("Error renaming file: %s\n", err)
		} else {
			m.Files[i] = newName
		}
	}
}

func (m *Map) updateFiles() {
	if m.Active {
		m.enableFiles()
	} else {
		m.disableFiles()
	}
}

func (m *Map) toggleActive() {
	m.Active = !m.Active
}

func (m *Map) addPath(path string) {
	if len(m.Files) == 0 {
		m.Files = append(m.Files, path)
	} else {
		exists := false
		for i := 0; i < len(m.Files); i++ {
			exists = m.Files[i] == path
		}
		if !exists {
			m.Files = append(m.Files, path)
		}
	}
}

type Game struct {
	Maps	[]*Map
}

func (g *Game) MapsFromPaths(items []string) {
	for i := 0; i < len(mapNames); i++ {
		for j := 0; j < len(items); j++ {
			if strings.Contains(items[j], mapNames[i]) {
				if len(g.Maps) == 0 {
					g.Maps = append(g.Maps, &Map{Name: mapNames[i], Active:true, Files:[]string{items[j]}})
				} else {
					exists := true
					for k := 0; k < len(g.Maps); k++ {
						if g.Maps[k].Name == mapNames[i] {
							g.Maps[k].Files = append(g.Maps[k].Files, items[j])
							exists = true
						}
					}
					if !exists {
						g.Maps = append(g.Maps, &Map{Name: mapNames[i], Active:true, Files:[]string{items[j]}})
					}
				}
			}
		}
	}
}

func (g *Game) ActiveFromFilenames() {
	for i := 0; i < len(g.Maps); i++ {
		g.Maps[i].activeFromFiles()
	}
}

func (g *Game) FilenamesFromActive() {
	for i := 0; i < len(g.Maps); i++ {
		g.Maps[i].updateFiles()
	}
}