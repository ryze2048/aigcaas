package aigcaas

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type AsyncHandlerData struct {
	ApplicationName       string
	ApiName               string
	Version               string
	CommonText2ImgRequest *CommonText2ImgRequest
	CommonImg2ImgRequest  *CommonImg2ImgRequest
}

// AsyncHandler
// 适用于大多数Stable Diffusion模型
// 包含文生图｜图生图
// 只支持异步调用 支持轮询模式和通知模式
func (c *Client) AsyncHandler(data *AsyncHandlerData) (requestId string, err error) {
	var response *http.Response
	var requestUrl = c.GetUrl(data.ApplicationName, data.ApiName, data.Version)
	switch data.ApiName {
	case TEXT2IMG: // 文生图
		if response, err = c.Do(data.CommonText2ImgRequest, requestUrl); err != nil {
			return "", err
		}
	case IMG2IMG: // 图生图
		if response, err = c.Do(data.CommonImg2ImgRequest, requestUrl); err != nil {
			return "", err
		}
	default:
		return "", ParameterError
	}
	defer func() {
		_ = response.Body.Close()
	}()
	var body = make([]byte, 0)
	if body, err = io.ReadAll(response.Body); err != nil {
		return "", err
	}
	switch response.Header.Get(`Aigcaas-Status`) {
	case SuccessStatus:
		var commonSuccessResponse CommonSuccessResponse
		if err = json.Unmarshal(body, &commonSuccessResponse); err != nil {
			return "", err
		}
		return commonSuccessResponse.AigcaasRequestId, err
	case ErrorsStatus:
		var commonErrorResponse CommonErrorResponse
		if err = json.Unmarshal(body, &commonErrorResponse); err != nil {
			return "", err
		}
		return "", errors.New(commonErrorResponse.Message)
	default:
		return "", errors.New(string(body))
	}
}
