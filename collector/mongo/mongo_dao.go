package mongo

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"cmpService/collector/collectdevice"
	"cmpService/common/lib"
)

type MongoAccessor struct {
	Address string
	Database string
	Table string
	Session *mgo.Session
	Collection *mgo.Collection
}

var Mongo *MongoAccessor

func NewMongoAccessor(address, db, c string) *MongoAccessor {
	session, err := mgo.Dial(address)
	if err != nil {
		fmt.Println("ERROR: failed to create mongodb connection!!")
		return nil
	}
	collection := session.DB(db).C(c)
	return &MongoAccessor{
		Address : address,
		Database : db,
		Table : c,
		Session :    session,
		Collection : collection,
	}
}

func SetMongo(mongo *MongoAccessor) {
	Mongo = mongo
	lib.LogWarn("Mongo IP:%s DB:%s TABLE:%s\n",
		mongo.Address, mongo.Database, mongo.Table)
}

func (m *MongoAccessor) Close() error {
	m.Session.Close()
	return nil
}

func idToObjectId(id collectdevice.ID) bson.ObjectId {
	return bson.ObjectIdHex(string(id))
}

func objectIdToID(objID bson.ObjectId) collectdevice.ID {
	return collectdevice.ID(objID)
}

func (m *MongoAccessor) Get(id collectdevice.ID) (collectdevice.ColletDevice, error) {
	t := collectdevice.ColletDevice{}
	err := m.Collection.FindId(idToObjectId(id)).One(&t)
	return t, err
}

func (m *MongoAccessor) Put(id collectdevice.ID, d collectdevice.ColletDevice) error {
	return m.Collection.UpdateId(idToObjectId(id), d)
}

func (m *MongoAccessor) Post(d *collectdevice.ColletDevice) (collectdevice.ID, error) {
	objID := bson.NewObjectId()
	d.Id = collectdevice.ID(fmt.Sprintf("%x",string(objID)))
	_, err := m.Collection.UpsertId(objID, &d)
	return d.Id, err
}

func (m *MongoAccessor) DeleteAll() (*mgo.ChangeInfo, error) {
	return m.Collection.RemoveAll(nil)
}

func (m *MongoAccessor) Delete(id collectdevice.ID) error {
	return m.Collection.RemoveId(idToObjectId(id))
}

func (m *MongoAccessor) GetAll() ([]collectdevice.ColletDevice, error) {
	var devices []collectdevice.ColletDevice
	err := m.Collection.Find(nil).All(&devices)
	return devices, err
}


