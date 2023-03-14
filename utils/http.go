package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/yangkaiyue/gin-exp/global"
	"io"
	"net/http"
)

type HttpCli struct {
	*http.Client
	Method string
	URL    string
	Body   []byte
	Header map[string]string
}

// Do 根据参数新建请求
func (cli *HttpCli) Do() (content []byte, err error) {

	if cli.Client == nil {
		cli.Client = http.DefaultClient
	}

	// Create Request
	req, err := http.NewRequest(cli.Method, cli.URL, bytes.NewReader(cli.Body))
	if err != nil {
		global.Logger.Error("Create Http Request Failed. Err: ", err.Error())
		return
	}

	// Set header
	req.Header.Add("Content-Type", "application/json")
	for k, v := range cli.Header {
		req.Header.Add(k, v)
	}

	// Send && defer Close
	resp, err := cli.Client.Do(req)
	if err != nil {
		global.Logger.Error("Send HTTP Request Failed. Err: ", err.Error())
		return
	}
	defer resp.Body.Close()

	// Parse Response
	content, err = io.ReadAll(resp.Body)
	if err != nil {
		global.Logger.Error("Parse HTTP Response Failed. Err: ", err.Error())
		return
	}

	// Check Response Error Code
	if resp.StatusCode != 200 {
		global.Logger.Error("HTTP Status Code NOT 200. Err: ", err.Error())
		return
	}

	return
}

// GetToken 请求 Global 获取 Token, 不使用配置文件内容而是传参是为了支持不同调用
func (cli *HttpCli) GetToken(URL, user, password string) (token string, err error) {

	if cli.Client == nil {
		cli.Client = http.DefaultClient
	}

	data, _ := json.Marshal(map[string]string{
		"account":  user,
		"password": password,
	})
	uri := fmt.Sprintf("%v/api/global-platform/sys/user/userLogin", URL)
	httpCli := HttpCli{
		Client: cli.Client,
		Method: "POST",
		URL:    uri,
		Body:   data,
	}
	content, err := httpCli.Do()
	if err != nil {
		return
	}

	// 由于 Response Error Message 格式不同, 单独处理
	if gjson.GetBytes(content, "head.errCode").Int() != 0 {
		err = errors.New(gjson.GetBytes(content, "head.errMsg").String())
		global.Logger.Error("Error Code NOT 0. Err: ", err.Error())
		return
	}

	// 返回 token
	token = gjson.GetBytes(content, "data").String()
	return
}
