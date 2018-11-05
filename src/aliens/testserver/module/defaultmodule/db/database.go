package db

import (
	"aliens/aliensbot/database/mongo"
	"aliens/testserver/module/commonmodule/conf"
)

var Database *mongo.Database = &mongo.Database{}

func Init() {
	err := Database.Init(conf.Config.Database)
	if err != nil {
		panic(err)
	}
	//Database.EnsureTable("collection", &protocol.Collection{})
}

func Close() {
	Database.Close()
}