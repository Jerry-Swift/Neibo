package poc

import "github.com/imroc/req/v3"

type Poc interface {
	SendPoc(target string, info map[string]interface{})
	SaveResult(target string, output string)
	CheckExp(resp *req.Response, target string, info map[string]interface{}) bool
}
