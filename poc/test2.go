package poc

import (
	"Corgi/common"
	"Corgi/configs"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Test2 struct{}

func (test2 Test2) SendPoc(target string, urlInfo configs.UrlInfo) {
	logrus.Info("[+] Start Test")
	urlInfo.Url = target
	urlInfo.Method = "Get"
	resp := common.Req(urlInfo)
	fmt.Println("resp:", resp)
}
