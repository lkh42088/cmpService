package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetMyPublicIp() string {
	url := "https://domains.google.com/checkip"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("GetMyPublicIp: error", err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("GetMyPublicIp: error", err)
	}
	fmt.Println("response:", string(data))
	return string(data)
}

