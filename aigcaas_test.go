package aigcaas

import (
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
	var client = NewClient(``, ``)
	var requestUrl = client.GetUrl(Alladin3Alladin3, TEXT2IMG, Version2)
	commonText2ImgRequest := &CommonText2ImgRequest{Prompt: "1girl,death,reaper,weapon"}
	var err error
	var response *http.Response
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
