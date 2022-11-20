package egnethttpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

var httpTr = &http.Transport{
	//控制主机的最大空闲连接数，0为没有限制
	MaxIdleConns:        100,
	MaxIdleConnsPerHost: 100,
	//长连接在关闭之前，保持空闲的最长时间，0表示没限制。
	IdleConnTimeout: 60 * time.Second,
}

var client = &http.Client{
	Timeout:   20 * time.Second,
	Transport: httpTr,
}

func post() {

	var err error

	var (
		method string
		url    string
		// param    io.Reader
		respData interface{}
	)

	method = "post"
	url = "http://127.0.0.1:8080/hello"

	// reqParam := map[string]string{
	// 	"str1": "ok",
	// }

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("str1", "aaaaa")

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		err = fmt.Errorf("create http request failed: %s", err.Error())
		// return err
	}

	// writer.FormDataContentType()

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("send http request failed: %s", err.Error())
		// return err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("read response data failed: %s", err.Error())
		// return err
	}
	// logger.Debug("read the response ok:", string(respBytes))

	// 解析响应
	if err = json.Unmarshal(respBytes, respData); err != nil {
		// logger.Error(err.Error())
		err = fmt.Errorf("unmarshal response data failed: %s", err.Error())
		// return err
	}

	fmt.Printf("respData: %v", respData)
	// return nil

	// resp, err := http.PostForm("http://127.0.0.1:8080/hello",
	// 	url.Values{"key": {"Value"}, "id": {"123"}})

	// fmt.Printf("resp: %v, err: %v", resp, err)
}
