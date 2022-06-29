package poc

import (
	"Corgi/common"
	"Corgi/configs"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Test struct{}

func (test Test) SendPoc(target string, pocInfo configs.UrlInfo) {
	logrus.Info("[+] Start Test")
	pocInfo.Url = target
	pocInfo.Method = "Get"
	resp := common.Req(pocInfo)
	fmt.Println("resp:", resp)

}
