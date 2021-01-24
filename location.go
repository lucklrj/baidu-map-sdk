package baidu_map_sdk

import netUrl "net/url"

func (c *Client) GetLngAndLatByAddress(address string) string {
	url := "/geocoding/v3/"
	params := netUrl.Values{}
	params.Add("address", address)
	return c.call(url, params)
}
