package mcapi

import (
	"bytes"
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/svcmgr/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendMcRegisterServer(server mcmodel.McServerDetail) bool {
	fmt.Printf("McServer : %v\n", server)
	pbytes, _ := json.Marshal(server)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlRegisterServer)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendAddVm: error 1 ", err)
		return false
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendAddVm: error 2 ", err)
		return false
	}
	fmt.Println("response: ", string(data))

	var mcserver mcmodel.MgoServer
	json.Unmarshal(data, &mcserver)
	// Dao: Server
	fmt.Println("mcserver:", mcserver.Mac)
	s := server.McServer
	s.Mac = mcserver.Mac
	s.Status = 1
	s, err = config.SvcmgrGlobalConfig.Mariadb.UpdateMcServer(s)
	if err != nil {
		fmt.Println("UpdateMcServer: error - ", err)
	}

	// Dao: Images
	fmt.Println("images:", mcserver.Images)
	if mcserver.Images != nil {
		for _, img := range *mcserver.Images {
			var image mcmodel.McImages
			image.Name = img.Name
			image.Hdd = img.Hdd
			image.Variant = img.Variant
			image.McServerIdx = int(s.Idx)
			image, err = config.SvcmgrGlobalConfig.Mariadb.AddMcImage(image)
			fmt.Println("insert image: ", image)
		}
	}

	// Dao: Networks
	fmt.Println("networks:", mcserver.Networks)
	if mcserver.Networks != nil {
		for _, net := range *mcserver.Networks {
			var network mcmodel.McNetworks
			network.Name = net.Name
			network.Bridge = net.Bridge
			network.Mode = net.Mode
			network.Ip = net.Ip
			network.Netmask= net.Netmask
			network.Prefix = net.Prefix
			network.McServerIdx = int(s.Idx)
			network, err = config.SvcmgrGlobalConfig.Mariadb.AddMcNetwork(network)
			fmt.Println("insert network: ", network)
		}
	}

	return true
}

func SendMcUnRegisterServer(server mcmodel.McServer) bool {
	pbytes, _ := json.Marshal(server)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlUnRegisterServer)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendAddVm: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendAddVm: error 2 ", err)
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}

func SendAddVm(vm mcmodel.McVm, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(vm)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlCreateVm)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendAddVm: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendAddVm: error 2 ", err)
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}

func SendDeleteVm(vm mcmodel.McVm, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(vm)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlDeleteVm)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendDeleteVm: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendDeleteVm: error 2 ", err)
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}

func SendGetVmById(vm mcmodel.McVm, server mcmodel.McServerDetail) bool {
	url := fmt.Sprintf("http://%s:8082%s%s/%d",server.IpAddr, lib.McUrlPrefix, lib.McUrlVm, vm.Idx)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("SendGetVmById: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendGetVmById: error 2 ", err)
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}

func SendGetVmAll(server mcmodel.McServerDetail) bool {
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlVm)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("SendGetVmAll: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendGetVmAll: error 2 ", err)
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}

func SendAddNetwork(net mcmodel.McNetworks, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(net)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",
		server.IpAddr, lib.McUrlPrefix, lib.McUrlNetworkAdd)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendAddVm: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendAddVm: error 2 ", err)
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}

func SendDeleteNetwork(net mcmodel.McNetworks, server mcmodel.McServerDetail) bool {
	pbytes, _ := json.Marshal(net)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:8082%s%s",
		server.IpAddr, lib.McUrlPrefix, lib.McUrlNetworkDelete)
	response, err := http.Post(url, "application/json", buff)
	if err != nil {
		fmt.Println("SendDeleteVm: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendDeleteVm: error 2 ", err)
		return false
	}
	fmt.Println("response: ", string(data))
	return true
}
