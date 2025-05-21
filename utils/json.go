package utils

import (
	"encoding/json"
	"github.com/dean0731/dean-tool/exception"
)

func JsonDump(value interface{}) string {
	v, err := json.Marshal(value)
	if err != nil {
		panic(exception.JsonDumpError)
	}
	return string(v)
}

func JsonLoad(date []byte) interface{} {
	var result interface{}
	err := json.Unmarshal(date, &result)
	if err != nil {
		panic(exception.JsonLoadError.SetMessageArgs(string(date)))
	}
	return result
}
