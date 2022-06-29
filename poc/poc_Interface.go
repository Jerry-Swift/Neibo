package poc

import (
	"Corgi/configs"
)

type Poc interface {
	SendPoc(target string, urlInfo configs.UrlInfo)
	//SaveResult(target string, output string)
	//CheckExp(resp *req.Response, target string, info map[string]interface{}) bool
}
