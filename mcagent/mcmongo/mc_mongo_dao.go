package mcmongo

import (
	"cmpService/common/mcmodel"
	"github.com/globalsign/mgo"
)

func (m *McMongoAccessor) AddVm(vm *mcmodel.MgoVm) (id int, err error) {
	id = int(vm.Idx)
	_, err = m.Collection.UpsertId(id, vm)
	return id, err
}

func (m*McMongoAccessor) GetVmAll() (vms []mcmodel.MgoVm, err error) {
	err = m.Collection.Find(nil).All(&vms)
	return vms, err
}

func (m*McMongoAccessor) GetVmById(id int) (vm mcmodel.MgoVm, err error) {
	err = m.Collection.FindId(id).One(&vm)
	return vm, err
}

func (m *McMongoAccessor) DeleteVm(id int) error {
	return m.Collection.RemoveId(id)
}

func (m *McMongoAccessor) DeleteVmAll() (*mgo.ChangeInfo, error) {
	return m.Collection.RemoveAll(nil)
}
