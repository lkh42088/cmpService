package rest

import (
	"bytes"
	"cmpService/common/messages"
	"cmpService/common/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRegisterUser (t *testing.T) {
	url := restServer + "/api/auth/register"
	userMsg := messages.UserRegisterMessage{
		Id:                 "nubes",
		Password:           "nubes1510",
		Email:              "bhjung@nubes-bridge.com",
		Name:               "정병화",
		EmailAuthFlag:      false,
		EmailAuthGroupFlag: false,
		EmailAuthGroupList: nil,
	}

	pbytes, _ := json.Marshal(userMsg)
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

func TestRegisterUserAuth (t *testing.T) {
	url := restServer + "/api/auth/register"
	userMsg := messages.UserRegisterMessage{
		Id:                 "nubes2",
		Password:           "nubes1510",
		Email:              "bhjung@nubes-bridge.com",
		Name:               "정병화",
		EmailAuthFlag:      true,
		EmailAuthGroupFlag: false,
		EmailAuthGroupList: nil,
	}

	pbytes, _ := json.Marshal(userMsg)
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

func TestRegisterUserAuthGroup (t *testing.T) {
	url := restServer + "/api/auth/register"
	userMsg := messages.UserRegisterMessage{
		Id:                 "nubes3",
		Password:           "nubes1510",
		Email:              "bhjung@nubes-bridge.com",
		Name:               "정병화",
		EmailAuthFlag:      false,
		EmailAuthGroupFlag: true,
		//EmailAuthGroupList: []string{"bhjung@nubes-bridge.com", "byeonghwa.jung@gmail.com"},
	}

	pbytes, _ := json.Marshal(userMsg)
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

func TestUnregisterUser(t *testing.T) {
	url := restServer + "/api/auth/unregister"
	userMsg := messages.UserRegisterMessage{
		Id:                 "nubes2",
		Email:              "",
		Name:               "",
		EmailAuthFlag:      false,
		EmailAuthGroupFlag: false,
		EmailAuthGroupList: nil,
	}
	pbytes, _ := json.Marshal(userMsg)
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
		//ID:"andrew",
		//Password: "andrew1510",
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
