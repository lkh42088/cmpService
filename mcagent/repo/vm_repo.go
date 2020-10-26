package repo

import (
	"cmpService/common/mcmodel"
	"cmpService/mcagent/config"
	"fmt"
)

var GlobalVmCache []mcmodel.McVm

func GetVmCacheObject() []mcmodel.McVm {
	return GlobalVmCache
}

func GetVmCache() *[]mcmodel.McVm {
	return &GlobalVmCache
}

func getVmByName(name string) (int, *mcmodel.McVm){
	for index, vm := range GlobalVmCache {
		if vm.Name == name {
			return index, &vm
		}
	}
	return -1, nil
}

func UpdateVmList(list *[]mcmodel.McVm) {
	if list == nil {
		return
	}

	fmt.Println("UpdateVmList --- ")
	for _, obj := range *list {
		_, vm := getVmByName(obj.Name)
		if vm != nil {
			if vm.Compare(obj) == true {
				fmt.Println("UpdateVmList: change vm!", vm.Name)
				// Changed vm
				UpdateVm2Repo(&obj)
			}
		}
	}
}

func AddVm2Repo(v *mcmodel.McVm) bool {
	if v.Name == "" {
		return false
	}
	for _, obj := range GlobalVmCache {
		if obj.Name == v.Name {
			return false
		}
	}
	// Get Available Vm Index
	v.VmIndex = config.GetAvailableVmIndex()
	// Allocate Vm Index
	config.AllocateVmIndex(uint(v.VmIndex))
	// Get Micro Cloud Server Idx
	v.McServerIdx = int(GetMcServer().Idx)
	// Insert to Cache List
	GlobalVmCache = append(GlobalVmCache, *v)

	/********************
	 * Add to Database
	 ********************/
	AddVm2Db(*v)
	return true
}

func AddVm2RepoForSync(v *mcmodel.McVm) bool {
	if v.Name == "" {
		return false
	}
	for _, obj := range GlobalVmCache {
		if obj.Name == v.Name {
			return false
		}
	}
	// Allocate Vm Index
	config.AllocateVmIndex(uint(v.VmIndex))
	// Insert to Cache List
	GlobalVmCache = append(GlobalVmCache, *v)
	/********************
	 * Add to Database
	 ********************/
	AddVm2Db(*v)
	return true
}

func UpdateVm2Repo(v *mcmodel.McVm) bool {
	if v.Name == "" {
		return false
	}

	for i, obj := range GlobalVmCache {
		if obj.Name == v.Name {
			/*************************
			 * Update from Database
			 *************************/
			dbVm, err := GetVmFromDbByName(v.Name)
			if err != nil {
				return false
			}
			GlobalVmCache[i] = dbVm.Update(*v)
			fmt.Println("UpdateVm2Repo: update ", dbVm.Name)
			UpdateVm2Db(GlobalVmCache[i])
			return true
		}
	}
	return false
}

func DeleteVmFromRepo(v mcmodel.McVm) bool {
	if v.Name == "" {
		return false
	}
	for i, obj := range GlobalVmCache {
		if obj.Name == v.Name {
			// Remove from Cache List
			GlobalVmCache = append(GlobalVmCache[:i], GlobalVmCache[i+1:]...)
			// Release Vm Index
			config.ReleaseVmIndex(uint(v.VmIndex))
			/*************************
			 * Delete from Database
			 *************************/
			dbVm, err  := GetVmFromDbByName(v.Name) // For real vm.Idx value
			if err == nil {
				DeleteVmFromDb(dbVm)
			}
			return true
		}
	}
	return false
}

func GetVmFromRepoByName(name string) *mcmodel.McVm {
	if name == "" {
		return nil
	}
	for _, obj := range GlobalVmCache {
		if obj.Name == name {
			return &obj
		}
	}
	vm, err := GetVmFromDbByName(name)
	if err == nil {
		// Caching vm
		GlobalVmCache = append(GlobalVmCache, vm)
		return &vm
	}
	return nil
}

func InitCachingVms() {
	/*************************
	 * Get Vm List from Database
	 *************************/
	list, err := GetAllVmFromDb()
	if err != nil {
		return
	}
	for _, obj := range list {
		GlobalVmCache = append(GlobalVmCache, obj)
		// Apply vm index to Global config
		config.AllocateVmIndex(uint(obj.VmIndex))
	}
}

func AddVm2Db(v mcmodel.McVm) (vm mcmodel.McVm, err error) {
	return config.GetMcGlobalConfig().DbOrm.AddMcVm(v)
}

func UpdateVm2Db(v mcmodel.McVm) (vm mcmodel.McVm, err error) {
	return config.GetMcGlobalConfig().DbOrm.UpdateMcVm(v)
}

func DeleteVmFromDb(v mcmodel.McVm) (vm mcmodel.McVm, err error) {
	return config.GetMcGlobalConfig().DbOrm.DeleteMcVm(v)
}

func GetVmFromDbByName(name string) (mcmodel.McVm, error){
	return config.GetMcGlobalConfig().DbOrm.GetMcVmByName(name)
}

func GetAllVmFromDb() ([]mcmodel.McVm, error){
	return config.GetMcGlobalConfig().DbOrm.GetAllMcVm()
}

//backup
func AddBackup2Db(v mcmodel.McVmBackup) (mcmodel.McVmBackup, error) {
	return config.GetMcGlobalConfig().DbOrm.AddMcVmBackup(v)
}