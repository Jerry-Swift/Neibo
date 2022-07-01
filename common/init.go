package common

import "Corgi/configs"

func UrlInit() configs.UrlInfo { // Url配置初始化
	urlInfo := configs.UrlInfo{
		Url:     "",
		UrlPath: "",
		Headers: make(map[string]string),
		Method:  "",
		Retry:   3,
	}
	return urlInfo
}

func InfoInit() configs.GlobalInfo { // 全局配置初始化
	info := configs.GlobalInfo{
		IP:             "",
		Port:           "",
		Mode:           "",
		Pocs:           "",
		OutPath:        "",
		Timeout:        3,
		Threads:        100,
		Specific_Ports: "",
		Proxy:          "",
	}
	return info
}
