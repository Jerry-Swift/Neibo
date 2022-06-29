package common

import (
	"Corgi/configs"
	"fmt"
	"github.com/imroc/req/v3"
)

func Init() configs.UrlInfo {
	pocInfo := configs.UrlInfo{
		Url:     "",
		UrlPath: "",
		Headers: make(map[string]string),
		Method:  "",
		Retry:   3,
		Body:    make(map[string]string),
		Proxy:   "",
		Timeout: 5,
	}
	return pocInfo
}
func Req(info configs.UrlInfo) *req.Response {
	client := req.C()
	if info.Method == "Get" {
		resp, err := client.R().
			SetBody(info.Body).
			SetHeaders(info.Headers).
			Get(info.Url + info.UrlPath)

		if err != nil {
			//fmt.Println("Request Url Wrong:", info.Url+info.UrlPath)
			panic(err)
		}
		return resp

	} else {
		resp, err := client.R().
			SetBody(info.Body).
			SetHeaders(info.Headers).
			Post(info.Url + info.UrlPath)
		if err != nil {
			fmt.Println("Request Url Wrong:", info.Url+info.UrlPath)
		}
		return resp
	}
}
