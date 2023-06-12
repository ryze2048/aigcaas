package aigcaas

// URL 请求地址
const URL string = `https://api.aigcaas.cn`

const AsyncRequestIdURL = `https://api.aigcaas.cn/v2/async`

// HeaderStatus 响应头状态参数
const HeaderStatus string = `Aigcaas-Status`

// Version2 版本
const (
	Version2 = `v2`
)

// 类型
const (
	TEXT2IMG        = `text2img` // 文生图
	IMG2IMG  string = `img2img`  // 图生图
)

// 异步请求的模式
const (
	_                int = iota
	PollingMode          // 轮询模式
	NotificationMode     // 通知模式
)

// API响应头Aigcaas-Status参数
const (
	SuccessStatus string = `True`  // 成功
	ErrorsStatus  string = `False` // 失败
)

const (
	ResponseSuccessCode int = 200 // 成功
	ResponseWaitCode        = 202 // 等待
)
