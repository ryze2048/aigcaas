package aigcaas

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// SetHeader
// 设置请求头｜计算签名
// 支持同步｜异步调用（支持轮询模式和通知模式）
func (c *Client) SetHeader(request *http.Request) {
	var timestamp = time.Now().Unix()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	nonce := rand.Intn(10000)
	sign := strconv.FormatInt(timestamp, 10) + c.SecretKey + strconv.Itoa(nonce)
	hash := sha256.New()
	hash.Write([]byte(sign))
	var token = hex.EncodeToString(hash.Sum(nil))
	request.Header.Add("SecretID", c.SecretId)
	request.Header.Add("Nonce", strconv.Itoa(nonce))
	request.Header.Add("Token", token)
	request.Header.Add("Timestamp", strconv.FormatInt(timestamp, 10))
	request.Header.Add("Content-Type", "application/json")
	switch c.Mode {
	case PollingMode: // 轮询模式
		request.Header.Add("Aigcaas-Async", "True")
	case NotificationMode: // 通知模式
		request.Header.Add("Aigcaas-Async", "True")
		request.Header.Add("Aigcaas-Async-Callback", c.Callback)
	}
}
