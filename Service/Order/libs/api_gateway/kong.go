package api_gateway

import (
	"gopkg.in/resty.v1"
	"fmt"
	"go-module/config/api_gateway"
)
	
func RegisterKong() {
	fmt.Println("=======START KONG=======")
	client := resty.New()
	res, _ := client.R().	
				SetFormData(map[string]string{
			"name": api_gateway.SERVICE_NAME,
			"path": api_gateway.SERVICE_PATH,
			"url": api_gateway.HOST_ADDRESS,
		}).Post(api_gateway.GATEWAY_ADDRESS)
	fmt.Println(res)
}
