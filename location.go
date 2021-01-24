package baidu_map_sdk

import "net/url"

func (c *Client) GetLngAndLatByAddress(address string) string {
	uri := "/geocoding/v3/"
	params := url.Values{}
	params.Add("address", address)
	return c.call(uri, params)
}
