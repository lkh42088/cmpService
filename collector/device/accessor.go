package device

type Accessor interface {
	Get(id ID) (Device, error)
	Put(id ID, d Device) error
	Post(d Device) (ID, error)
	Delete(id ID) error
}


