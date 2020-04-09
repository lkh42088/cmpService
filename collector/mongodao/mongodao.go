package mongodao

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"nubes/collector/device"
)

type MongoAccessor struct {
	session *mgo.Session
	collection *mgo.Collection
}

//var Mongo = device.NewMemoryDataAccess()
var Mongo = New("127.0.0.1", "collector", "devices")

func New(path, db, c string) *MongoAccessor {
	session, err := mgo.Dial(path)
	if err != nil {
		return nil
	}
	collection := session.DB(db).C(c)
	return &MongoAccessor{
		session:    session,
		collection: collection,
	}
}

func (m *MongoAccessor) Close() error {
	m.session.Close()
	return nil
}

func idToObjectId(id device.ID) bson.ObjectId {
	return bson.ObjectIdHex(string(id))
}

func objectIdToID(objID bson.ObjectId) device.ID {
	return device.ID(objID)
}

func (m *MongoAccessor) Get(id device.ID) (device.Device, error) {
	t := device.Device{}
	err := m.collection.FindId(idToObjectId(id)).One(&t)
	return t, err
}

func (m *MongoAccessor) Put(id device.ID, d device.Device) error {
	return m.collection.UpdateId(idToObjectId(id), d)
}

func (m *MongoAccessor) Post(d *device.Device) (device.ID, error) {
	objID := bson.NewObjectId()
	d.Id = device.ID(fmt.Sprintf("%x",string(objID)))
	_, err := m.collection.UpsertId(objID, &d)
	return d.Id, err
}

func (m *MongoAccessor) DeleteAll() (*mgo.ChangeInfo, error) {
	return m.collection.RemoveAll(nil)
}

func (m *MongoAccessor) Delete(id device.ID) error {
	return m.collection.RemoveId(idToObjectId(id))
}

func (m *MongoAccessor) GetAll() ([]device.Device, error) {
	var devices []device.Device
	err := m.collection.Find(nil).All(&devices)
	return devices, err
}


