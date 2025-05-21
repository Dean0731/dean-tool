package log

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

func Debug(msg ...interface{}) {
	if viper.GetBool("D") {
		Println(msg...)
	}
}

func Debugf(msg string, args ...interface{}) {
	Debug(fmt.Sprintf(msg, args...))
}

func Println(msg ...interface{}) {
	fmt.Println(msg...)
}

func Printf(msg string, args ...interface{}) {
	Println(fmt.Sprintf(msg, args...))
}

func PrintJson(jsonStr string) {
	if !viper.GetBool("J") {
		Println(jsonStr)
		return
	}
	var jsonData interface{}
	if err := json.Unmarshal([]byte(jsonStr), &jsonData); err != nil {
		panic(err)
	}
	var formattedJSON []byte
	formattedJSON, err := json.MarshalIndent(jsonData.(map[string]interface{})["Response"], "", "  ") // 使用两个空格作为缩进
	if err != nil {
		panic(err)
	}
	Println(string(formattedJSON))
}
