package common

import (
	"bytes"
	"github.com/imroc/req/v3"
	"math/rand"
	"time"
)

//TODO
//  为了解决添加随机路由功能，需要加入生成随机字符串函数 √
// 各pocs在发送请求时，要能将body传送给Req函数

func Req(reqOptions map[string]interface{}) *req.Response {
	// 为了加快扫描速度，每次发请求时都开新内存，存储请求信息
	url := reqOptions["url"].(string)
	method := reqOptions["method"].(string)
	headers := reqOptions["headers"].(map[string]string)
	body := reqOptions["body"].(map[string]string)
	retry := reqOptions["retry"].(int)
	timeout := reqOptions["timeout"].(int)
	proxy := reqOptions["proxy"].(string)
	redirect := reqOptions["redirect"].(bool)

	client := req.C()
	req := client
	//老代码，太过冗余
	//client := req.C()
	//if info.Method == "Get" {
	//	resp, err := client.R().
	//		SetHeaders(info.Headers).
	//		Get(info.Url + info.UrlPath)
	//
	//	if err != nil {
	//		//fmt.Println("Request Url Wrong:", info.Url+info.UrlPath)
	//		panic(err)
	//	}
	//	return resp
	//
	//} else {
	//	resp, err := client.R().
	//		SetBody(info.Body).
	//		SetHeaders(info.Headers).
	//		Post(info.Url + info.UrlPath)
	//	if err != nil {
	//		fmt.Println("Request Url Wrong:", info.Url+info.UrlPath)
	//	}
	//	return resp
	//}
}

// 生成指定长度的随机字符串
func RandomString(codeLen int) string {
	rawStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	buf := make([]byte, 0, codeLen) // 开一块长度为len的内存，创建长度为0的切片buf
	buffer := bytes.NewBuffer(buf)
	rand.Seed(time.Now().UnixNano()) //精确到纳秒
	for rawStrLen := len(rawStr); codeLen > 0; codeLen-- {
		randNum := rand.Intn(rawStrLen)
		buffer.WriteByte(rawStr[randNum])
	}

	return buffer.String()
}
