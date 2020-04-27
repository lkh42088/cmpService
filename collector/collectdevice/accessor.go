package collectdevice

type Accessor interface {
	Get(id ID) (ColletDevice, error)
	Put(id ID, d ColletDevice) error
	Post(d ColletDevice) (ID, error)
	Delete(id ID) error
}


