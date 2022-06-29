package poc

import (
	"Corgi/common"
	"Corgi/configs"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Test2 struct{}

func (test2 Test2) SendPoc(target string, urlInfo configs.UrlInfo) {
	logrus.Info("[+] Start Test2")
	urlInfo.Url = target
	urlInfo.Method = "Post"
	urlInfo.Body["Bo"] = "Test Body"
	resp := common.Req(urlInfo)
	fmt.Println("resp:", resp)
}

//
//func (test2 Test2) CheckExp(target string, output string) {
//
//}
//
//func (test2 Test2) SaveResult(resp *req.Response, target string, info map[string]interface{}) bool {
//	return true
//}
