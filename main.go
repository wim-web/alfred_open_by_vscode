package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"search_files/alfred"
	"search_files/setting"
	"strings"
)

func main() {

	args := os.Args

	searchWord := args[1]
	settingPath := args[2]

	folders := setting.Read(settingPath)

	scriptFilters := gatherFiles(folders, searchWord)

	b, _ := json.Marshal(alfred.Output{Items: scriptFilters})

	fmt.Print(string(b))
}

func gatherFiles(folders setting.Folders, searchWord string) (scriptFilters []alfred.ScriptFilter) {

	for _, folder := range folders.Folders {
		files, err := ioutil.ReadDir(folder.Path)
		if err != nil {
			o := alfred.Output{}
			o.Error(err.Error())
			os.Exit(1)
		}

		for _, file := range files {

			if searchWord == ":" {
				scriptFilters = append(scriptFilters, alfred.ScriptFilter{
					Type:     "file",
					Title:    file.Name(),
					SubTitle: filepath.Join(folder.Path, file.Name()),
					Arg:      filepath.Join(folder.Path, file.Name()),
					Valid:    true,
				})
				continue
			}

			if strings.Contains(file.Name(), searchWord) {
				scriptFilters = append(scriptFilters, alfred.ScriptFilter{
					Type:     "file",
					Title:    file.Name(),
					SubTitle: filepath.Join(folder.Path, file.Name()),
					Arg:      filepath.Join(folder.Path, file.Name()),
					Valid:    true,
				})
			}
		}
	}

	return
}
