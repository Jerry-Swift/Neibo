package poc

import (
	"Corgi/configs"
)

type Poc interface {
	SendPoc(target string, urlInfo configs.UrlInfo)
	//SaveResult(target string, output string) //TODO 加入结果输出和验证函数
	//CheckExp(resp *req.Response, target string, info map[string]interface{}) bool
}
