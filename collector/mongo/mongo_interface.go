package mongo

import (
	"github.com/globalsign/mgo"
	"cmpService/collector/collectdevice"
)

type MongoDBLayer interface {
	Get(collectdevice.ID) (collectdevice.ColletDevice, error)
	Put(collectdevice.ID, collectdevice.ColletDevice) error
	Post(*collectdevice.ColletDevice) (collectdevice.ID, error)
	DeleteAll() (*mgo.ChangeInfo, error)
	Delete(collectdevice.ID) error
	GetAll() ([]collectdevice.ColletDevice, error)
}
