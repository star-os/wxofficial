package util

import (
	"bytes"
	json "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

// ErrInfo 是微信平台返回的通用错误信息定义
// 用于避免重复定义
type ErrInfo struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// HttpPostJSON 用于发送携带JSON数据的HTTP POST请求
func HttpPostJSON(url string, obj interface{}) ([]byte, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	// 可能要考虑JSON中转义字符的问题
	body := bytes.NewReader(data)

	resp, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// HttpGet 用于发送HTTP GET请求
func HttpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body) // Go1.16版本可以改为`io.ReadAll`
}
