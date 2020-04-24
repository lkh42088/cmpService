package mongodao

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"nubes/collector/device"
)

type MongoAccessor struct {
	Session *mgo.Session
	Collection *mgo.Collection
}

//var Mongo = device.NewMemoryDataAccess()
//var Mongo = New("127.0.0.1", "collector", "devices")
var Mongo *MongoAccessor

func New(path, db, c string) *MongoAccessor {
	session, err := mgo.Dial(path)
	if err != nil {
		fmt.Println("ERROR: failed to create mongodb connection!!")
		return nil
	}
	collection := session.DB(db).C(c)
	return &MongoAccessor{
		Session:    session,
		Collection: collection,
	}
}

func SetMongo(mongo *MongoAccessor) {
	Mongo = mongo
}

func (m *MongoAccessor) Close() error {
	m.Session.Close()
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
	err := m.Collection.FindId(idToObjectId(id)).One(&t)
	return t, err
}

func (m *MongoAccessor) Put(id device.ID, d device.Device) error {
	return m.Collection.UpdateId(idToObjectId(id), d)
}

func (m *MongoAccessor) Post(d *device.Device) (device.ID, error) {
	objID := bson.NewObjectId()
	d.Id = device.ID(fmt.Sprintf("%x",string(objID)))
	_, err := m.Collection.UpsertId(objID, &d)
	return d.Id, err
}

func (m *MongoAccessor) DeleteAll() (*mgo.ChangeInfo, error) {
	return m.Collection.RemoveAll(nil)
}

func (m *MongoAccessor) Delete(id device.ID) error {
	return m.Collection.RemoveId(idToObjectId(id))
}

func (m *MongoAccessor) GetAll() ([]device.Device, error) {
	var devices []device.Device
	err := m.Collection.Find(nil).All(&devices)
	return devices, err
}


