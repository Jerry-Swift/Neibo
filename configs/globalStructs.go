package configs

type GlobalInfo struct {
	IP             string
	Port           string
	TargetsFile    string
	Mode           string // 扫描速度模式
	Pocs           string // 测试的pocs
	OutPath        string
	Timeout        int
	Threads        int
	Specific_Ports string // 默认探测的端口
	Proxy          string
	Verbose        int // 输出结果的详细等级
	UrlInfo
}

type UrlInfo struct {
	Url     string
	UrlPath string
	Headers map[string]string
	Body    map[string]string
	Method  string
	Retry   int
}
