package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"cmpService/common/models"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	url := restServer + "/v1/register"
	user := models.User{
		ID:"andrew",
		Password: "andrew1510",
		Email: "andrew@cmpService-bridge.com",
		Name:"anrew",
		Level: 1,
	}
	pbytes, _ := json.Marshal(user)
	buff := bytes.NewBuffer(pbytes)

	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("error 1: ", err)
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error 2: ", err)
		return
	}
	fmt.Println("response:", string(data))
}

func TestLoginUser(t *testing.T) {
	url := restServer + "/login"
	user := models.User{
		ID:"andrew",
		Password: "andrew1510",
	}
	pbytes, _ := json.Marshal(user)
	buff := bytes.NewBuffer(pbytes)

	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("error 1: ", err)
		return
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error 2: ", err)
		return
	}
	fmt.Println("response:", string(data))
}
