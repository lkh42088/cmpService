package mcmongo

import (
	"cmpService/common/lib"
	"cmpService/mcagent/config"
	"fmt"
	"github.com/globalsign/mgo"
)

type McMongoAccessor struct {
	Address string
	Database string
	Table string
	Session *mgo.Session
	Collection *mgo.Collection
}

var McMongo *McMongoAccessor

func NewMcMongoAccessor(address, db, c string) *McMongoAccessor {
	session, err := mgo.Dial(address)
	if err != nil {
		fmt.Println("ERROR: failed to create mongodb connection!!")
		return nil
	}
	collection := session.DB(db).C(c)
	return &McMongoAccessor{
		Address:    address,
		Database:   db,
		Table:      c,
		Session:    session,
		Collection: collection,
	}
}

func SetMcMongo(mongo *McMongoAccessor) {
	McMongo = mongo
	lib.LogWarn("Mongo IP:%s DB:%s TABLE:%s\n",
		mongo.Address, mongo.Database, mongo.Table)
}

func (m *McMongoAccessor) Close() error {
	m.Session.Close()
	return nil
}

func Configure() bool {
	cfg := config.GetMcGlobalConfig()
	if cfg.MongoIp == "" || cfg.MongoDb == "" || cfg.MongoCollection == "" {
		lib.LogWarn("Failed MongoDb configuration!\n")
		return false
	}
	// Configure MongoDB
	m := NewMcMongoAccessor(cfg.MongoIp, cfg.MongoDb, cfg.MongoCollection)
	if m != nil {
		SetMcMongo(m)
		return true
	}
	return false
}