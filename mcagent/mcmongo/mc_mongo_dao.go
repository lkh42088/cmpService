package mcmongo

import (
	"cmpService/common/mcmodel"
	"github.com/globalsign/mgo"
)
//
//func (m *McMongoAccessor) AddVm(vm *mcmodel.MgoVm) (id int, err error) {
//	id = int(vm.Idx)
//	_, err = m.Collection.UpsertId(id, vm)
//	return id, err
//}
//
//func (m *McMongoAccessor) UpdateVmByInternal(vm *mcmodel.MgoVm) (id int, err error) {
//	fmt.Println("UPDATE: ", vm)
//	//err = m.Collection.UpdateId(vm.Idx, mcmodel.MgoVm{
//	//	Filename: vm.Filename,
//	//	IpAddr: vm.IpAddr,
//	//	Mac: vm.Mac,
//	//	CurrentStatus: vm.CurrentStatus,
//	//	VmNumber: vm.VmNumber,
//	//})
//	//return int(vm.Idx), err
//	return m.AddVm(vm)
//}
//
//func (m*McMongoAccessor) GetVmAll() (vms []mcmodel.MgoVm, err error) {
//	err = m.Collection.Find(nil).All(&vms)
//	return vms, err
//}
//
//func (m*McMongoAccessor) GetVmById(id int) (vm mcmodel.MgoVm, err error) {
//	err = m.Collection.FindId(id).One(&vm)
//	return vm, err
//}

func (m *McMongoAccessor) DeleteVm(id int) error {
	return m.Collection.RemoveId(id)
}

func (m *McMongoAccessor) DeleteVmAll() (*mgo.ChangeInfo, error) {
	return m.Collection.RemoveAll(nil)
}

// Image
func (m *McMongoAccessor) AddImage(obj *mcmodel.MgoImage) (err error) {
	err = m.Session.DB(m.Database).C("vm_image").UpdateId(obj.Id, obj)
	return err
}

func (m *McMongoAccessor) DeleteImage(id int) error {
	return m.Session.DB(m.Database).C("vm_image").RemoveId(id)
}

func (m*McMongoAccessor) GetImageAll() (images []mcmodel.MgoImage, err error) {
	err = m.Collection.Find(nil).All(&images)
	return images, err
}
