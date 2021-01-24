package main

import (
	"fmt"

	baiduSDK "github.com/lucklrj/baidu-map-sdk"
)

func main() {
	ak := "123"
	sk := "456"
	client := baiduSDK.New(ak, sk, true)

	address := "百度大厦"
	result := client.GetLngAndLatByAddress(address)
	fmt.Println(result)
}
