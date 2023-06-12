package aigcaas

import "fmt"

// GetStableDiffusionUrl 基于模型 SD 1.5 请求的API地址
func (c *Client) GetStableDiffusionUrl(applicationName, apiName string) string {
	return fmt.Sprintf("%s/v2/application/%s/api/%s", URL, applicationName, apiName)
}
