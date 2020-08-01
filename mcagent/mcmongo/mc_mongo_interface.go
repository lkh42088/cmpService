package mcmongo

import "cmpService/mcagent/mcmodel"

type McMongoDBLayer interface {
	AddVm(vm *mcmodel.VmEntry) (id int, err error)
	GetVmById(id int) (vm mcmodel.VmEntry, err error)
}