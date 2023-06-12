package aigcaas

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestAsyncHandler(t *testing.T) {
	var client = NewClient(``, ``)
	asyncHandlerData := AsyncHandlerData{
		ApplicationName:       Alladin3Alladin3,
		ApiName:               TEXT2IMG,
		Version:               Version2,
		CommonText2ImgRequest: &CommonText2ImgRequest{Prompt: "1girl,death,reaper,weapon"},
	}
	var requestId string
	var err error
	if requestId, err = client.AsyncHandler(&asyncHandlerData); err != nil {
		t.Error("async handler err --> ", err)
	}
	t.Log("request id --> ", requestId)
	return
}

func TestAigcaas(t *testing.T) {
	var err error
	var response *http.Response
	var client = NewClient(``, ``)
	var requestUrl = client.GetUrl(Alladin3Alladin3, TEXT2IMG, Version2)
	commonText2ImgRequest := &CommonText2ImgRequest{Prompt: "1girl,death,reaper,weapon"}
	if response, err = client.Do(commonText2ImgRequest, requestUrl); err != nil {
		t.Error("client do err --> ", err)
	}
	defer func() {
		_ = response.Body.Close()
	}()
	var body = make([]byte, 0)
	if body, err = io.ReadAll(response.Body); err != nil {
		t.Error("io read all --> ", err)
	}

	t.Log("body info --> ", string(body))
}

func TestClient_AsyncRequestId(t *testing.T) {
	var err error
	var requestId = `e94e43c0-c888-46a0-a972-2062a215f501`
	// var requestId = ``
	var client = NewClient(``, ``)
	var response *http.Response
	if response, err = client.AsyncRequestId(requestId); err != nil {
		t.Error("client async request id err --> ", err)
	}
	defer func() {
		_ = response.Body.Close()
	}()
	var bodyByte = make([]byte, 0)
	if bodyByte, err = io.ReadAll(response.Body); err != nil {
		t.Error("io read all err --> ", err)
	}
	t.Log("body info --> ", string(bodyByte))
	switch response.StatusCode {
	case ResponseSuccessCode:
		// 成功
		var asyncRequestIdInfoSuccessResponse AsyncRequestIdInfoSuccessResponse
		if err = json.Unmarshal(bodyByte, &asyncRequestIdInfoSuccessResponse); err != nil {
			t.Error("json unmarshal err --> ", err)
		}
		fmt.Println(asyncRequestIdInfoSuccessResponse)
	case ResponseWaitCode:
		// 等待
		var asyncRequestIdWaitResponse AsyncRequestIdWaitResponse
		if err = json.Unmarshal(bodyByte, &asyncRequestIdWaitResponse); err != nil {
			t.Error("json unmarshal err --> ", err)
		}
		fmt.Println(asyncRequestIdWaitResponse)
	default:
		// 失败
		var asyncRequestIdErrorResponse AsyncRequestIdErrorResponse
		if err = json.Unmarshal(bodyByte, &asyncRequestIdErrorResponse); err != nil {
			t.Error("json unmarshal err --> ", err)
		}
		fmt.Println(asyncRequestIdErrorResponse)
	}
}
