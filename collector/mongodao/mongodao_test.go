package mongodao

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"nubes/collector/device"
	"testing"
)

func TestMongoGetAll(t *testing.T) {
	devices, _ := Mongo.GetAll()
	for i, dev := range devices {
		fmt.Println("Device ", dev, i)
	}
}

func TestMongoPost(t *testing.T) {
	objId := bson.NewObjectId()
	fmt.Println("objId:", objId)
	d := device.Device{
		Id:            device.ID(objId),
		Ip:            "192.168.10.115",
		Port:          161,
		SnmpCommunity: "nubes",
	}
	fmt.Println("device:", d)
	devId, err := Mongo.Post(&d)
	fmt.Println("id:", fmt.Sprintf("%x",devId), err)
	fmt.Println(bson.ObjectId(devId), err)
}

func TestMongoDelete(t *testing.T) {
	s := "5e79c9a4902ed2796f1878d0"
	Mongo.Delete(device.ID(s))
}

func TestMongoDeleteAll(t *testing.T) {
	desc, _ := Mongo.DeleteAll()
	fmt.Println(desc)
}

