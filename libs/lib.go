package libs

import (
	"path/filepath"
	"os"
	"fmt"
	"strings"
)

var mapNames = [2]string{"desert", "erangel"}

type MapFile struct {
	active	bool
	path	string
}

func (m MapFile) getActive() bool {
	return m.active
}

func (m *MapFile) disableFile() {
	newName := m.path + ".disabled"
	err := os.Rename(m.path, newName)
	if err != nil {
		fmt.Printf("Error renaming file: %s\n", err)
	} else {
		m.path = newName
	}
}

func (m *MapFile) enableFile() {
	newName := m.path[:len(m.path)-len(".disabled")]
	err := os.Rename(m.path, newName)
	if err != nil {
		fmt.Printf("Error renaming file: %s\n", err)
	} else {
		m.path = newName
	}
}

func (m *MapFile) toggleActive() {
	if m.active {
		m.disableFile()
	} else {
		m.enableFile()
	}
	m.active = !m.active
}

type Map struct {
	name   string
	active bool
	files  []string
}

func (m *Map) addFile(file string) {
	if len(m.files) == 0 {
		m.files = append(m.files, file)
	} else {
		exists := false
		for i := 0; i < len(m.files); i++ {
			exists = m.files[i] == file
		}
		if !exists {
			m.files = append(m.files, file)
		}
	}
}

func (m *Map) updateActive() {
	m.active = filepath.Ext(m.files[0]) != ".disabled"
}

func (m *Map) enableFiles() {
	for i := 0; i < len(m.files); i++ {
		newName := m.files[i][:len(m.files[i])-len(".disabled")]
		err := os.Rename(m.files[i], newName)
		if err != nil {
			fmt.Printf("Error renaming file: %s\n", err)
		} else {
			m.files[i] = newName
		}
	}
}

func (m *Map) disableFiles() {
	for i := 0; i < len(m.files); i++ {
		newName := m.files[i] + ".disabled"
		err := os.Rename(m.files[i], newName)
		if err != nil {
			fmt.Printf("Error renaming file: %s\n", err)
		} else {
			m.files[i] = newName
		}
	}
}

func (m *Map) updateFiles() {
	if m.active {
		m.enableFiles()
	} else {
		m.disableFiles()
	}
}

func (m *Map) toggleActive() {
	m.active = !m.active
	m.updateFiles()
}

/** Map End **/

type Game struct {
	Maps	[]*Map
}

func (g *Game) MapsFromPaths(items []string) {
	for i := 0; i < len(mapNames); i++ {
		for j := 0; j < len(items); j++ {
			if strings.Contains(items[j], mapNames[i]) {
				if len(g.Maps) == 0 {
					g.Maps = append(g.Maps, &Map{name: mapNames[i], active:true, files:[]string{items[j]}})
				} else {
					exists := false
					for k := 0; k < len(g.Maps); k++ {
						if g.Maps[k].name == mapNames[i] {
							g.Maps[k].files = append(g.Maps[k].files, items[j])
							exists = true
						}
					}
					if !exists {
						g.Maps = append(g.Maps, &Map{name: mapNames[i], active:true, files:[]string{items[j]}})
					}
				}
			}
		}
	}
}

func (g Game) UpdateMaps() {
	for i := 0; i < len(g.Maps); i++ {
		g.Maps[i].updateActive()
	}
}

func (g Game) UpdateFiles() {
	for i := 0; i < len(g.Maps); i++ {
		g.Maps[i].updateFiles()
	}
}

func (g Game) ToggleActive(id int) {
	g.Maps[id].toggleActive()
}

func (g Game) GetMaps() map[string]bool {
	maps := make(map[string]bool)

	for i := 0; i < len(g.Maps); i++ {
		maps[g.Maps[i].name] = g.Maps[i].active
	}

	return maps
}

/** Game end **/