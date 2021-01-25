package baidu_map_sdk

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/tidwall/gjson"
)

//通过地址获取经纬度
func (c *Client) GetLngAndLatByAddress(address string) (float64, float64) {
	uri := "/geocoding/v3/"
	params := url.Values{}
	params.Add("address", address)
	response := c.call(uri, params)
	lng := gjson.Parse(response).Get("result.location.lng").Float()
	lat := gjson.Parse(response).Get("result.location.lat").Float()
	return lat, lng
}

//根据经纬度获取地址
func (c *Client) GetAddressByLngAndLat(lng float64, lat float64) string {
	uri := "/reverse_geocoding/v3/"
	params := url.Values{}
	params.Add("location", strconv.FormatFloat(lat, 'g', 1, 64)+","+strconv.FormatFloat(lng, 'g', 1, 64))
	response := c.call(uri, params)
	fmt.Println(response)
	return ""
}
