package baidu_map_sdk

import (
	"crypto/md5"
	"fmt"
	"net/url"

	"github.com/ddliu/go-httpclient"
)

type Client struct {
	ak         string
	sk         string
	httpClient *httpclient.HttpClient
	isNeedSn   bool
}

func (c *Client) call(uri string, data url.Values) string {
	data.Add("output", "json")
	data.Add("ak", c.ak)
	if c.isNeedSn == true {
		data.Add("sn", getSn(uri, data, c.sk))
	}

	//格式转换
	response, _ := c.httpClient.Begin().Get("http://api.map.baidu.com"+uri, data)
	responseString, _ := response.ToString()
	return responseString
}

func New(ak string, sk string, isNeedSn bool) *Client {
	httpClient := httpclient.NewHttpClient().Defaults(httpclient.Map{
		httpclient.OPT_USERAGENT: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:64. 0) Gecko/20100101 Firefox/64.0",
		//httpclient.OPT_USERAGENT:  "Mozilla/5.0 (Windows NT 6.1; rv:24.0) Gecko/20100101 Firefox/24.0",
		httpclient.OPT_UNSAFE_TLS: true,
	})
	return &Client{
		ak:         ak,
		sk:         sk,
		httpClient: httpClient,
		isNeedSn:   isNeedSn,
	}
}
func getSn(uri string, data url.Values, sk string) string {
	//todo post的时候需要按key排序
	query := data.Encode()
	return fmt.Sprintf("%x", md5.Sum([]byte(url.QueryEscape(uri+"?"+query+sk))))
}
