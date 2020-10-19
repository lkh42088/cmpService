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

func ApplyMcServerResource(recvMsg mcmodel.McServerMsg, server mcmodel.McServerDetail) {
	// Dao: Server
	fmt.Println("recvMsg:", recvMsg.Mac)
	recvMsg.Dump()
	s := server.McServer
	s.Mac = recvMsg.Mac
	s.Status = 1
	s.Port = recvMsg.Port
	s.IpAddr = recvMsg.Ip
	s.PublicIpAddr = recvMsg.PublicIp
	// backup data
	s.UcloudAccessKey = recvMsg.UcloudAccessKey
	s.UcloudSecretKey = recvMsg.UcloudSecretKey
	s.UcloudProjectId = recvMsg.UcloudProjectId
	s.UcloudDomainId = recvMsg.UcloudDomainId
	s.NasUrl = recvMsg.NasUrl
	s.NasId = s.NasId
	s.NasPassword = recvMsg.NasPassword
	s, err := config.SvcmgrGlobalConfig.Mariadb.UpdateMcServer(s)
	if err != nil {
		fmt.Println("UpdateMcServer: error - ", err)
	}

	// Dao: Images
	fmt.Println("images:", recvMsg.Images)
	if recvMsg.Images != nil {
		imgList, _ := config.SvcmgrGlobalConfig.Mariadb.GetMcImagesByServerIdx(int(s.Idx))
		for _, img := range *recvMsg.Images {
			if mcmodel.LookupImage(&imgList, img) != nil {
				continue
			}
			img.McServerIdx = int(s.Idx)
			obj, _ := config.SvcmgrGlobalConfig.Mariadb.AddMcImage(img)
			fmt.Println("insert image: ", obj)
		}
	}

	// Dao: Networks
	fmt.Println("networks:", recvMsg.Networks)
	if recvMsg.Networks != nil {
		netList, _ := config.SvcmgrGlobalConfig.Mariadb.GetMcNetworksByServerIdx(int(s.Idx))
		for _, net := range *recvMsg.Networks {
			dbnet := mcmodel.LookupNetwork(&netList, net)
			if dbnet != nil {
				net.Idx = dbnet.Idx
				net.McServerIdx = dbnet.McServerIdx
				obj, _ := config.SvcmgrGlobalConfig.Mariadb.UpdateMcNetwork(net)
				fmt.Println("update network: ", obj)
			} else {
				net.Idx = 0
				net.McServerIdx = int(s.Idx)
				obj, _ := config.SvcmgrGlobalConfig.Mariadb.AddMcNetwork(net)
				fmt.Println("insert network: ", obj)
			}
		}
	}

	if recvMsg.Vms != nil {
		vmList, _ := config.SvcmgrGlobalConfig.Mariadb.GetMcVmsByServerIdx(int(s.Idx))
		for _, vm := range *recvMsg.Vms {
			old := mcmodel.LookupVm(&vmList, vm)
			if old != nil {
				vm.Idx = old.Idx
				vm.CompanyIdx = old.CompanyIdx
				vm.McServerIdx = old.McServerIdx
				obj, _ := config.SvcmgrGlobalConfig.Mariadb.UpdateMcVm(vm)
				fmt.Println("update vm: ", obj)
			} else {
				vm.Idx = 0
				vm.McServerIdx = int(s.Idx)
				vm.CompanyIdx = s.CompanyIdx
				obj, _ := config.SvcmgrGlobalConfig.Mariadb.AddMcVm(vm)
				fmt.Println("insert vm: ", obj)
			}
		}
		for _, vm := range vmList {
			obj := mcmodel.LookupVm(&vmList, vm)
			if obj == nil {
				config.SvcmgrGlobalConfig.Mariadb.DeleteMcVm(vm)
			}
		}
	} else {
		vmList, _ := config.SvcmgrGlobalConfig.Mariadb.GetMcVmsByServerIdx(int(s.Idx))
		for _, vm := range vmList {
			config.SvcmgrGlobalConfig.Mariadb.DeleteMcVm(vm)
		}
	}
}

func SendMcRegisterServer(server mcmodel.McServerDetail) bool {
	fmt.Printf("McServer : %v\n", server)
	pbytes, _ := json.Marshal(server)
	buff := bytes.NewBuffer(pbytes)
	url := fmt.Sprintf("http://%s:%s%s%s",
		server.IpAddr, server.L4Port,
		lib.McUrlPrefix, lib.McUrlRegisterServer)
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

	/** process msg ***/
	var mcserver mcmodel.McServerMsg
	json.Unmarshal(data, &mcserver)

	/** process msg ***/
	ApplyMcServerResource(mcserver, server)

	return true
}

func SendMcRegisterServerOld(server mcmodel.McServerDetail) bool {
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

	var mcserver mcmodel.McServerMsg
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

func SendGetMcServer(server mcmodel.McServerDetail) bool {
	url := fmt.Sprintf("http://%s:8082%s%s",server.IpAddr, lib.McUrlPrefix, lib.McUrlMonServer)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("SendGetMcServerAll: error 1 ", err)
		return false
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("SendGetMcServerAll: error 2 ", err)
		return false
	}
	fmt.Println("response: ", string(data))
	var serv mcmodel.McServerMsg
	err = json.Unmarshal(data, &serv)
	if err != nil {
		fmt.Println("SendGetMcServerAll: error 3 ", err)
		return false
	}
	serv.Dump()
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
