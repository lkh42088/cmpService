package agent

import (
	"cmpService/common/mcmodel"
)

type VmList struct {
	List map[uint]mcmodel.MgoVm
}

var McVms *VmList

func NewVmList() *VmList {
	return &VmList{
		List: map[uint]mcmodel.MgoVm{},
	}
}

func SetVmList(list *VmList) {
	McVms = list
}

func ConfigureVmList() {
	n := NewVmList()
	SetVmList(n)
}

func (v *VmList)GetVm(id uint) (*mcmodel.MgoVm, bool) {
	vm, exist := v.List[id]
	if !exist {
		return nil, exist
	}
	return &vm, exist
}

func (v *VmList)AddVm(vm mcmodel.MgoVm) {
	v.List[vm.Idx] = vm
}

func (v *VmList)UpdateVm(vm mcmodel.MgoVm) {
	v.List[vm.Idx] = vm
}

func (v *VmList)DeleteVm(id uint) {
	if _, exists := v.List[id]; !exists {
		return
	}
	delete(v.List, id)
}
