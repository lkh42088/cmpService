package mariadblayer

import (
	"cmpService/common/mcmodel"
)

func (db *DBORM) GetSysPlatform() (platform []mcmodel.DevicePlatform, err error) {
	err = db.
		Table("sysinfo_tb").
		Select("cpu_model, COUNT(cpu_model) as count").
		Group("cpu_model").
		Find(&platform).Error

	return platform, err
}

func (db *DBORM) GetVmOsInfo() (osInfo []mcmodel.DeviceOsInfo, err error) {
	err = db.
		Table("mc_vm_tb").
		Select("vm_os, COUNT(vm_os) as count").
		Group("vm_os").
		Find(&osInfo).Error

	return osInfo, err
}


