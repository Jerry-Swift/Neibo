package attack

import (
	"Corgi/common"
	"Corgi/poc"
)

// 拿到url后默认自动扫描所有poc，也可指定poc名称

func Soldier(url string) {

	pocInfo := common.Init()
	exp := poc.Test{}
	exp.SendPoc(url, pocInfo)
}
