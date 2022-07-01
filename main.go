package main

import (
	"Corgi/attack"
	"Corgi/configs"
)

// TODO 加入命令行参数  √
//     日志文件功能
func main() {
	//target := "https://httpbin.org/get"
	configs.Banner()
	globalOptions := attack.ParseOptions()
	attack.Process(globalOptions)

}
