package utils

import (
	"fmt"
	"github.com/dean0731/dean-tool/exception"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func HttpGet(url string) string {
	// 发起GET请求
	resp, err := http.Get(url)
	if err != nil {
		panic(exception.HttpError.SetMessage(err.Error()))
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(exception.UnknownError.SetMessage(err.Error()))
	}

	// 打印返回的内容
	return string(body)
}

func HttpGetDownload(url string, dir string, name string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(exception.HttpDownloadError.SetMessage(err.Error()))
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(exception.HttpError.SetMessageArgs(fmt.Sprintf("status code: %d %s", resp.StatusCode, resp.Status)))
	}
	if name == "" {
		ext := filepath.Ext(strings.Split(url, "?")[0])
		name = "tmp-name" + ext
	}

	tmpfile, err := os.CreateTemp(dir, name)
	if err != nil {
		panic(exception.CreateFileError.SetMessageArgs(dir))
	}

	if _, err = io.Copy(tmpfile, resp.Body); err != nil {
		os.Remove(tmpfile.Name()) // 如果写入失败，尝试删除临时文件
		panic(exception.CopyFileError.SetMessageArgs(tmpfile.Name()))
	}
	return tmpfile.Name()
}

func HttpPost(url string, contentType string, body io.Reader) string {
	value := GetBody(body)
	body = strings.NewReader(value)
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		panic(exception.HttpError.SetMessage(err.Error()))
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(exception.UnknownError.SetMessage(err.Error()))
	}
	return strings.TrimSpace(string(respBody))
}

func ExtractLastPathSegment(inputURL string) string {
	// 解析传入的URL
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return ""
	}

	// 移除路径两端的斜杠并分割路径
	path := strings.Trim(parsedURL.Path, "/")
	segments := strings.Split(path, "/")

	// 检查是否有足够的段
	if len(segments) == 0 {
		return ""
	}

	// 返回最后一个段
	return segments[len(segments)-1]
}

func GetBody(body io.Reader) string {
	// 读取请求体
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		panic(fmt.Sprintf("Failed to read body: %v", err))
	}

	data := string(bodyBytes)
	return data
}

var re = regexp.MustCompile(`<[^>]+>`)

func RemoveHTMLTags(html string) string {
	text := re.ReplaceAllString(html, "")

	// 去除可能存在的HTML实体
	text = strings.ReplaceAll(text, "&nbsp;", " ")
	// 如果有更多的HTML实体需要转换，可以继续使用strings.ReplaceAll或者使用html.UnescapeString

	// 处理可能的多余空格
	text = strings.Join(strings.Fields(text), " ")

	return text
}
