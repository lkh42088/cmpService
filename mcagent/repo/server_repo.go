package repo

import (
	"cmpService/common/mcmodel"
	config2 "cmpService/mcagent/config"
	"fmt"
)

var GlobalServerRepo mcmodel.McServerDetail

func GetMcServer() *mcmodel.McServerDetail {
	if GlobalServerRepo.SerialNumber == "" {
		servers, err := GetServerFromDb()
		if err != nil {
			return nil
		}
		if len(servers) == 1 {
			GlobalServerRepo.McServer = servers[0]
			// Company Name : todo 2020sep28 by bhjung
		} else {
			fmt.Printf("GetMcServer: server len(%d) is not 1 !!!\n", len(servers))
			return nil
		}
	}

	return &GlobalServerRepo
}

func UpdateMcServer(s mcmodel.McServerDetail) bool {
	fmt.Println("AddServer2Repo: Update")
	s.Idx = GlobalServerRepo.Idx
	temp, err := UpdateServer2Db(s.McServer)
	if err != nil {
		return false
	}
	GlobalServerRepo = s
	GlobalServerRepo.Idx = temp.Idx
	return true
}

func AddServer2Repo(s *mcmodel.McServerDetail) bool {
	list, _ := GetServerFromDb()
	if len(list) > 0 {
		return false
	}

	GlobalServerRepo = *s
	fmt.Println("AddServer2Repo: Add SN -", GlobalServerRepo.SerialNumber, "Idx", GlobalServerRepo.Idx)
	temp, _ := AddServer2Db(s.McServer)
	GlobalServerRepo.Idx = temp.Idx

	return true
}

func DeleteServer2Repo() bool {
	if GlobalServerRepo.SerialNumber == "" {
		return false
	}

	DeleteServer2Db(GlobalServerRepo.McServer)

	var newServer mcmodel.McServerDetail
	newServer.SerialNumber = GlobalServerRepo.SerialNumber
	AddServer2Repo(&newServer)
	return true
}

func AddServer2Db(s mcmodel.McServer) (mcmodel.McServer, error) {
	fmt.Println("AddServer2Db: ", s)
	return config2.GetMcGlobalConfig().DbOrm.AddMcServer(s)
}

func UpdateServer2Db(s mcmodel.McServer) (mcmodel.McServer, error) {
	fmt.Println("UpdateServer2Db: ", s)
	return config2.GetMcGlobalConfig().DbOrm.UpdateMcServerAll(s)
}

func DeleteServer2Db(s mcmodel.McServer) (mcmodel.McServer, error) {
	fmt.Println("DeleteServer2Db: ", s)
	return config2.GetMcGlobalConfig().DbOrm.DeleteMcServer(s)
}

func GetServerFromDb() ([]mcmodel.McServer, error) {
	return config2.GetMcGlobalConfig().DbOrm.GetMcServer()
}