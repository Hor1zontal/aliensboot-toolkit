package db

import (
	"github.com/KylinHe/aliensboot/database/mongo"
	"e.coding.net/aliens/aliensboot_testserver/module/room/conf"
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
