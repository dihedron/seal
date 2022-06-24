package logging

import (
	"encoding/json"
	"runtime"

	"gopkg.in/yaml.v3"
)

func GetCallerFrame(skip int) runtime.Frame {
	pc := make([]uintptr, 15)
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame
	//fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
}

func ToJSON(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}

func ToPrettyJSON(v interface{}) string {
	data, _ := json.MarshalIndent(v, "", "  ")
	return string(data)
}

func ToYAML(v interface{}) string {
	data, _ := yaml.Marshal(v)
	return string(data)
}
