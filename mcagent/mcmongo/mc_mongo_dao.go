package mcmongo

import "cmpService/mcagent/mcmodel"

func (m *McMongoAccessor) AddVm(vm *mcmodel.VmEntry) (id int, err error) {
	id = vm.Idx
	_, err = m.Collection.UpsertId(id, vm)
	return id, err
}

func (m*McMongoAccessor) GetVmById(id int) (vm mcmodel.VmEntry, err error) {
	err = m.Collection.FindId(id).One(&vm)
	return vm, err
}
