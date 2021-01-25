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
	lat, lng := client.GetLngAndLatByAddress(address)
	fmt.Println(lat, lng)

	client.GetAddressByLngAndLat(lat, lng)

}
