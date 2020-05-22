package mongo

import (
	"cmpService/collector/collectdevice"
	"github.com/globalsign/mgo"
)

type MongoDBLayer interface {
	Get(collectdevice.ID) (collectdevice.ColletDevice, error)
	Put(collectdevice.ID, collectdevice.ColletDevice) error
	Post(*collectdevice.ColletDevice) (collectdevice.ID, error)
	DeleteAll() (*mgo.ChangeInfo, error)
	Delete(collectdevice.ID) error
	GetAll() ([]collectdevice.ColletDevice, error)
}
