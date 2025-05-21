package pan

import (
	"encoding/json"
	"strings"

	_const "github.com/dean0731/dean-tool/const"
	"github.com/dean0731/dean-tool/utils"
)

var checkQuarkUrl = "https://drive-h.quark.cn/1/clouddrive/share/sharepage/token"
var quarkDomain = "pan.quark.cn"

func GetResourceFromId(url string) Response {
	id := utils.ExtractLastPathSegment(url)
	if id == "" {
		return Response{}
	}

	authReq := AuthRequest{
		PwdID:    id,
		Passcode: "",
	}

	// 将结构体编码为JSON
	jsonData, _ := json.Marshal(authReq)
	result := utils.HttpPost(checkQuarkUrl, _const.HttpContentTypeApplicationJson, strings.NewReader(string(jsonData)))
	var response Response
	err := json.Unmarshal([]byte(result), &response)
	if err != nil {
		panic(err.Error())
	}
	return response
}

func CheckResourceVaild(url string) string {
	if strings.Contains(url, quarkDomain) {
		return GetResourceFromId(url).Message
	}
	return "未知"
}
