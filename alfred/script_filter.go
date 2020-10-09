package alfred

import (
	"encoding/json"
	"fmt"
)

type ScriptFilter struct {
	Uid          string `json:"uid"`
	Type         string `json:"type"`
	Title        string `json:"title"`
	SubTitle     string `json:"subtitle"`
	Arg          string `json:"arg"`
	Icon         `json:"icon"`
	Valid        bool   `json:"valid"`
	Match        string `json:"match"`
	AutoComplete string `json:"autocomplete"`
	Mods         string `json:"mods"`
	Text         string `json:"text"`
	QuickLookUrl string `json:"quicklookurl"`
}

type Icon struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type Output struct {
	Items []ScriptFilter `json:"items"`
}

func (o Output) Error(message string) {
	i := append(o.Items, ScriptFilter{Title: message})
	o.Items = i
	j, _ := json.Marshal(o)
	fmt.Print(string(j))
}
