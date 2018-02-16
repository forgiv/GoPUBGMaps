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

func (m *MapFile) updateFile() {
	if m.active {
		m.disableFile()
	} else {
		m.enableFile()
	}
}

func (m *MapFile) toggleActive() {
	m.active = !m.active
	m.updateFile()
}

/** MapFile end **/

type Map struct {
	name   		string
	active 		bool
	mapfiles  	[]*MapFile
}

func (m *Map) addFile(file string) {
	if len(m.mapfiles) == 0 {
		m.mapfiles = append(m.mapfiles, &MapFile{active:filepath.Ext(file) != ".disabled", path:file})
	} else {
		exists := false
		for i := 0; i < len(m.mapfiles); i++ {
			exists = m.mapfiles[i].path == file
		}
		if !exists {
			m.mapfiles = append(m.mapfiles, &MapFile{active:filepath.Ext(file) != ".disabled", path:file})
		}
	}
}

func (m *Map) updateActive() {
	m.active = m.mapfiles[0].getActive()
}

func (m *Map) enableFiles() {
	for i := 0; i < len(m.mapfiles); i++ {
		m.mapfiles[i].enableFile()
	}
}

func (m *Map) disableFiles() {
	for i := 0; i < len(m.mapfiles); i++ {
		m.mapfiles[i].disableFile()
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
					g.Maps = append(g.Maps, &Map{name: mapNames[i], active:true, mapfiles:[]*MapFile{{active:filepath.Ext(items[j]) != ".disabled", path:items[j]}}})
				} else {
					exists := false
					for k := 0; k < len(g.Maps); k++ {
						if g.Maps[k].name == mapNames[i] {
							g.Maps[k].mapfiles = append(g.Maps[k].mapfiles, &MapFile{filepath.Ext(items[j]) != ".disabled", items[j]})
							exists = true
						}
					}
					if !exists {
						g.Maps = append(g.Maps, &Map{name: mapNames[i], active:true, mapfiles:[]*MapFile{{active:filepath.Ext(items[j]) != ".disabled", path:items[j]}}})
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