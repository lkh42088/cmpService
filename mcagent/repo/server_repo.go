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

func AddServer2Repo(s *mcmodel.McServerDetail) bool {
	if GlobalServerRepo.SerialNumber != "" && (
		GlobalServerRepo.SerialNumber != s.SerialNumber ||
		GlobalServerRepo.CompanyIdx != s.CompanyIdx) {
		temp, err := UpdateServer2Db(s.McServer)
		if err != nil {
			return false
		}
		GlobalServerRepo = *s
		GlobalServerRepo.Idx = temp.Idx
		return true
	}

	GlobalServerRepo = *s
	temp, _ := AddServer2Db(s.McServer)
	GlobalServerRepo.Idx = temp.Idx

	return true
}

func DeleteServer2Repo() bool {
	if GlobalServerRepo.SerialNumber == "" {
		return false
	}

	DeleteServer2Db(GlobalServerRepo.McServer)
	return true
}

func AddServer2Db(s mcmodel.McServer) (mcmodel.McServer, error) {
	return config2.GetMcGlobalConfig().DbOrm.AddMcServer(s)
}

func UpdateServer2Db(s mcmodel.McServer) (mcmodel.McServer, error) {
	return config2.GetMcGlobalConfig().DbOrm.UpdateMcServerAll(s)
}

func DeleteServer2Db(s mcmodel.McServer) (mcmodel.McServer, error) {
	return config2.GetMcGlobalConfig().DbOrm.DeleteMcServer(s)
}

func GetServerFromDb() ([]mcmodel.McServer, error) {
	return config2.GetMcGlobalConfig().DbOrm.GetMcServer()
}