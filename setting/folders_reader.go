package setting

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"search_files/alfred"
)

type Folders struct {
	Folders []Folder `json:"folders"`
}

type Folder struct {
	Path string `json:"path"`
}

func Read(filePath string) (f Folders) {

	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		o := alfred.Output{}
		o.Error(filePath + "が読み込めませんでした")
		os.Exit(1)
	}

	json.Unmarshal(raw, &f)

	return
}
