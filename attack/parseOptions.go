package attack

import (
	"Corgi/common"
	"Corgi/configs"
	"flag"
	"fmt"
)

func ParseOptions() *configs.GlobalInfo {

	globalInfo := &configs.GlobalInfo{}
	flag.StringVar(&globalInfo.IP, "i", "", "目标ip")
	flag.StringVar(&globalInfo.Port, "p", "8000", "目标端口") //TODO 需要添加扫描默认端口功能
	flag.StringVar(&globalInfo.Url, "u", "", "目标url")
	flag.StringVar(&globalInfo.TargetsFile, "f", "", "目标文件，支持ip与url，一行一个目标") //TODO 需要添加读取内容的处理代码，能自动分辨ip与url
	flag.StringVar(&globalInfo.Mode, "m", "crazy", "扫描的速度模式，默认为最速模式")
	flag.StringVar(&globalInfo.Pocs, "pocs", "", "验证的pocs，默认全部打一遍")
	flag.StringVar(&globalInfo.OutPath, "o", "", "输出文件，支持相对/绝对路径，默认在当前文件夹输出output.txt")
	flag.IntVar(&globalInfo.Timeout, "to", 3, "超时时间，默认3秒")
	flag.IntVar(&globalInfo.Retry, "r", 3, "请求失败重试次数")
	flag.StringVar(&globalInfo.Method, "method", "Get", "发送HTTP请求方式，默认Get")
	flag.IntVar(&globalInfo.Threads, "t", 100, "扫描线程数，默认100条") //TODO 添加协程
	flag.StringVar(&globalInfo.Proxy, "proxy", "", "代理服务器，格式:socks5://127.0.0.1:8888")
	flag.IntVar(&globalInfo.Verbose, "v", 0, "输出详细程度")
	flag.Parse()

	return globalInfo
}

// 流程处理函数，根据输入参数决定执行过程
func Process(globalOptions *configs.GlobalInfo) {
	if globalOptions.Url != "" {
		target := globalOptions.Url
		fmt.Println("Tartget URL:", target)
		Soldier(target, globalOptions)
	} else if globalOptions.Url == "" && globalOptions.TargetsFile == "" {
		fmt.Println("请输入目标")
	} else {
		//读取文件内目标
		urls, err := common.ReadFile(globalOptions.TargetsFile)
		//TODO 此处需要加入url处理代码

		if err != nil {
			fmt.Println("文件读取失败")
			panic(err)
		}
		for _, target := range urls {
			Soldier(target, globalOptions)
		}
	}
}
