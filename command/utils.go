package command

import (
	"encoding/json"

	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
)

var (
	red    = color.New(color.FgRed).SprintfFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
)

func ToJSON(value interface{}, pretty bool) string {
	var result []byte
	if pretty {
		result, _ = json.MarshalIndent(value, "", "  ")
	} else {
		result, _ = json.Marshal(value)
	}
	return string(result)
}

func ToYAML(value interface{}) string {
	result, _ := yaml.Marshal(value)
	return string(result)
}
