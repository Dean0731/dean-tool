package ocr

// This file is auto-generated, don't edit it. Thanks.

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	stream "github.com/alibabacloud-go/darabonba-stream/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	credential "github.com/aliyun/credentials-go/credentials"
	"github.com/mitchellh/mapstructure"
	"io"
)

func CreateClient(ak, sk string) (_result *openapi.Client, _err error) {
	// 工程代码建议使用更安全的无AK方式，凭据配置方式请参见：https://help.aliyun.com/document_detail/378661.html。
	c := new(credential.Config).SetType("access_key").SetAccessKeyId(ak).SetAccessKeySecret(sk)
	credential, _err := credential.NewCredential(c)
	if _err != nil {
		return _result, _err
	}

	config := &openapi.Config{
		Credential: credential,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/ocr-api
	config.Endpoint = tea.String("ocr-api.cn-hangzhou.aliyuncs.com")
	_result = &openapi.Client{}
	_result, _err = openapi.NewClient(config)
	return _result, _err
}

func CreateApiInfo() (_result *openapi.Params) {
	params := &openapi.Params{
		// 接口名称
		Action: tea.String("RecognizeAllText"),
		// 接口版本
		Version: tea.String("2021-07-07"),
		// 接口协议
		Protocol: tea.String("HTTPS"),
		// 接口 HTTP 方法
		Method:   tea.String("POST"),
		AuthType: tea.String("AK"),
		Style:    tea.String("V3"),
		// 接口 PATH
		Pathname: tea.String("/"),
		// 接口请求体内容格式
		ReqBodyType: tea.String("json"),
		// 接口响应体内容格式
		BodyType: tea.String("json"),
	}
	_result = params
	return _result
}

func OcrImage(ak, sk, filePath string, body io.Reader) (string, error) {
	client, _err := CreateClient(ak, sk)
	if _err != nil {
		return "", _err
	}

	params := CreateApiInfo()
	// query params
	queries := map[string]interface{}{}
	queries["Type"] = tea.String("Advanced")
	// 需要安装额外的依赖库，直接点击下载完整工程即可看到所有依赖。
	if body == nil {
		body = stream.ReadFromFilePath(tea.String(filePath))
	}
	// runtime options
	runtime := &util.RuntimeOptions{}
	request := &openapi.OpenApiRequest{
		Query:  openapiutil.Query(queries),
		Stream: body,
	}
	// 复制代码运行请自行打印 API 的返回值
	// 返回值实际为 Map 类型，可从 Map 中获得三类数据：响应体 body、响应头 headers、HTTP 返回的状态码 statusCode。
	result, _err := client.CallApi(params, request, runtime)
	if _err != nil {
		return "", _err
	} else {
		var ocr OCRResponse
		decoderConfig := &mapstructure.DecoderConfig{
			Result:  &ocr,
			TagName: "json", // 使用 json tag 来匹配字段
		}
		decoder, _ := mapstructure.NewDecoder(decoderConfig)
		_ = decoder.Decode(result["body"])
		return ocr.Data.Content, _err
	}
}

type OCRResponse struct {
	Data struct {
		Content       string `json:"Content"`
		Height        int    `json:"Height"`
		SubImageCount int    `json:"SubImageCount"`
		SubImages     []struct {
			Angle     int `json:"Angle"`
			BlockInfo struct {
				BlockCount   int `json:"BlockCount"`
				BlockDetails []struct {
					BlockAngle      int    `json:"BlockAngle"`
					BlockConfidence int    `json:"BlockConfidence"`
					BlockContent    string `json:"BlockContent"`
					BlockId         int    `json:"BlockId"`
				} `json:"BlockDetails"`
			} `json:"BlockInfo"`
			SubImageId int    `json:"SubImageId"`
			Type       string `json:"Type"`
		} `json:"SubImages"`
		Width int `json:"Width"`
	} `json:"Data"`
	RequestId string `json:"RequestId"`
}
