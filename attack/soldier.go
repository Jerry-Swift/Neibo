package attack

import (
	"Corgi/configs"
	"Corgi/poc"
	"github.com/fatih/structs"
	"strings"
)

//TODO:
//  拿到url后默认自动扫描所有poc √
// 也可指定poc名称
// 会将get和post都请求一遍，需要找到原因。
func Soldier(target string, info *configs.GlobalInfo) { // info为全局变量GlobalInfo
	pocInfo := info.UrlInfo
	pocInfo.Url = target

	pocs_init := make(map[string]interface{})
	pocs := pushPoc(pocs_init) // 直接创建字典需要make分配内存
	info_hashmap := structs.Map(&info)
	userPocs := info_hashmap["Pocs"].(string)
	pocsName := strings.Split(userPocs, ",")

	if len(pocsName) == 1 && pocsName[0] == "" {
		for _, poc_temp := range pocs {
			attack_poc := poc_temp.(poc.Poc)
			attack_poc.SendPoc(pocInfo.Url, pocInfo)
		}
	}

}

func pushPoc(pocs map[string]interface{}) map[string]interface{} {
	pocs["test"] = &poc.CVE202222947{}
	pocs["test2"] = &poc.Test2{}

	return pocs
}
