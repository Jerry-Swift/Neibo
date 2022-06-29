package configs

type GlobalInfo struct {
	IP             string
	Port           string
	Mode           string
	OutPath        string
	Timeout        int
	Threads        int
	Specific_Ports string
	Proxy          string
}

type UrlInfo struct {
	Url     string
	UrlPath string
	Headers map[string]string
	Method  string
	Retry   int
	Body    map[string]string
	Proxy   string
	Timeout int
}
