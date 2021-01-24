package baidu_map_sdk

import (
	"github.com/ddliu/go-httpclient"
)

type Client struct {
	ak string
	HttpClient *httpclient.HttpClient
	isNeedSn bool
}
func New(ak string,isNeedSn bool) *Client{
	httpClient := httpclient.NewHttpClient().Defaults(httpclient.Map{
		httpclient.OPT_USERAGENT: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:64. 0) Gecko/20100101 Firefox/64.0",
		//httpclient.OPT_USERAGENT:  "Mozilla/5.0 (Windows NT 6.1; rv:24.0) Gecko/20100101 Firefox/24.0",
		httpclient.OPT_UNSAFE_TLS: true,
	})
	return &Client{
		ak:ak,
		HttpClient: httpClient,
		isNeedSn:isNeedSn,
	}
}
