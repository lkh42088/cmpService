package mcmongo

type McMongoDBLayer interface {
	//AddVm(vm *mcmodel.MgoVm) (id int, err error)
	DeleteVm(id int) error
	//GetVmById(id int) (vm mcmodel.MgoVm, err error)
}