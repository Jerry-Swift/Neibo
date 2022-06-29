package attack

import (
	"Corgi/common"
	"Corgi/configs"
	"Corgi/poc"
	"github.com/fatih/structs"
	"strings"
)

// 拿到url后默认自动扫描所有poc，也可指定poc名称
// TODO:已实现自动填充pocs，但会将get和post都请求一遍，需要找到原因。
func Soldier(url string, info configs.GlobalInfo) { // info为全局变量GlobalInfo
	pocInfo := common.UrlInit()

	pocs_init := make(map[string]interface{})
	pocs := pushPoc(pocs_init) // 直接创建字典需要make分配内存
	info_hashmap := structs.Map(&info)
	userPocs := info_hashmap["Pocs"].(string)
	pocsName := strings.Split(userPocs, ",")

	if len(pocsName) == 1 && pocsName[0] == "" {
		for _, poc_temp := range pocs {
			attack_poc := poc_temp.(poc.Poc)
			attack_poc.SendPoc(url, pocInfo)
		}
	}

}

func pushPoc(pocs map[string]interface{}) map[string]interface{} {
	pocs["test"] = &poc.Test{}
	pocs["test2"] = &poc.Test2{}

	return pocs
}
