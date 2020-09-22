package agent

import (
	"cmpService/common/mcmodel"
)

type VmList struct {
	List map[uint]mcmodel.McVm
}

var McVms *VmList

func NewVmList() *VmList {
	return &VmList{
		List: map[uint]mcmodel.McVm{},
	}
}

func SetVmList(list *VmList) {
	McVms = list
}

func ConfigureVmList() {
	n := NewVmList()
	SetVmList(n)
}

//func GetVmListByDb() {
//	for _, vm := range McVms.List {
//		delete(McVms.List, vm.Idx)
//	}
//	vms, err := mcmongo.McMongo.GetVmAll()
//	if err != nil {
//		return
//	}
//	for _, vm := range vms {
//		McVms.List[vm.Idx] = vm
//	}
//}

func (v *VmList)GetVm(id uint) (*mcmodel.McVm, bool) {
	vm, exist := v.List[id]
	if !exist {
		return nil, exist
	}
	return &vm, exist
}

func (v *VmList)AddVm(vm mcmodel.McVm) {
	v.List[vm.Idx] = vm
}

func (v *VmList)UpdateVm(vm mcmodel.McVm) {
	v.List[vm.Idx] = vm
}

func (v *VmList)DeleteVm(id uint) {
	if _, exists := v.List[id]; !exists {
		return
	}
	delete(v.List, id)
}
