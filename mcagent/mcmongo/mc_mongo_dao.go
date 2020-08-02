package mcmongo

import "cmpService/common/mcmodel"

func (m *McMongoAccessor) AddVm(vm *mcmodel.MgoVm) (id int, err error) {
	id = int(vm.Idx)
	_, err = m.Collection.UpsertId(id, vm)
	return id, err
}

func (m*McMongoAccessor) GetVmById(id int) (vm mcmodel.MgoVm, err error) {
	err = m.Collection.FindId(id).One(&vm)
	return vm, err
}

func (m *McMongoAccessor) DeleteVm(id int) error {
	return m.Collection.RemoveId(id)
}
