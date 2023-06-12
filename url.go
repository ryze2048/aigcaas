package aigcaas

import "fmt"

func (c *Client) GetUrl(applicationName, apiName, version string) string {
	return fmt.Sprintf("%s/%s/application/%s/api/%s", URL, version, applicationName, apiName)
}
