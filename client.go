package aigcaas

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Client struct {
	SecretKey string `json:"secret_key"` // 密钥信息
	SecretId  string `json:"secret_id"`  // 密钥信息
	Mode      int    `json:"mode"`       // 模式
	Callback  string `json:"callback"`   // 异步通知模式时候需要填写回掉地址
}

var httpClient = http.Client{}

func NewClient(secretKey, secretId string) *Client {
	return &Client{
		SecretKey: secretKey,
		SecretId:  secretId,
	}
}

// Do 发送请求
// 注意：获取到response的处理结束后记得response.Body.Close()
func (c *Client) Do(body interface{}, url string) (response *http.Response, err error) {
	var request *http.Request
	var bodyByte = make([]byte, 0)
	if bodyByte, err = json.Marshal(body); err != nil {
		return nil, err
	}
	if request, err = http.NewRequest("POST", url, strings.NewReader(string(bodyByte))); err != nil {
		return nil, err
	}
	c.SetHeader(request)
	return httpClient.Do(request)
}

// AsyncRequestId 异步请求结果
func (c *Client) AsyncRequestId(requestId string) (response *http.Response, err error) {
	var request *http.Request
	if request, err = http.NewRequest("GET", AsyncRequestIdURL, nil); err != nil {
		return nil, err
	}
	c.SetHeader(request)
	request.Header.Add("RequestId", requestId)
	if response, err = httpClient.Do(request); err != nil {
		return nil, err
	}
	return
}
