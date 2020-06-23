package mongo

import (
	"cmpService/collector/collectdevice"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"testing"
)

func TestMongoGetAll(t *testing.T) {
	devices, _ := Mongo.GetAll()
	for i, dev := range devices {
		fmt.Println("ColletDevice ", dev, i)
	}
}

func TestMongoPost(t *testing.T) {
	objId := bson.NewObjectId()
	fmt.Println("objId:", objId)
	d := collectdevice.ColletDevice{
		Id:            collectdevice.ID(objId),
		Ip:            "192.168.10.115",
		Port:          161,
		SnmpCommunity: "cmpService",
	}
	fmt.Println("collectdevice:", d)
	devId, err := Mongo.Post(&d)
	fmt.Println("id:", fmt.Sprintf("%x", devId), err)
	fmt.Println(bson.ObjectId(devId), err)
}

func TestMongoDelete(t *testing.T) {
	s := "5e79c9a4902ed2796f1878d0"
	Mongo.Delete(collectdevice.ID(s))
}

func TestMongoDeleteAll(t *testing.T) {
	desc, _ := Mongo.DeleteAll()
	fmt.Println(desc)
}
