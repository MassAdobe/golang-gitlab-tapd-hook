package utils

import (
	"gitlab_tapd/common"
	"gitlab_tapd/logs"
	"io/ioutil"
	"net/http"
)

// GET请求封装
func Get(url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logs.MyError("GET Request NewRequest Failed", url, err.Error)
		panic(err.Error())
	}
	req.Header.Set(common.HTTP_HEADER_ENCODE_TITLE, common.HTTP_HEADER_ENCODE_INNER)
	req.Header.Set(common.HTTP_HEADER_PRIVATE_TOKEN_TITLE, common.HTTP_HEADER_PRIVATE_TOKEN_INNER)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.MyError("GET Request Failed", url, err.Error)
		panic(err.Error())
	}
	return body
}

// GetTapd
func TapdGet(url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logs.MyError("TapdGet Request NewRequest Failed", url, err.Error)
		panic(err.Error())
	}
	req.Header.Set(common.HTTP_HEADER_ENCODE_TITLE, common.HTTP_HEADER_ENCODE_INNER)
	req.Header.Set(common.HTTP_HEADER_AUTHORIZATION_TITLE, common.HTTP_HEADER_AUTHORIZATION_INNER)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.MyError("TapdGet Request Failed", url, err.Error())
		panic(err.Error())
	}
	return body
}
