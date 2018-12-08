package db

import (
	"aliens/aliensbot/database/mongo"
	"aliens/testserver/module/scene/conf"
)

var Database *mongo.Database = &mongo.Database{}

func Init() {
	err := Database.Init(conf.Config.Database)
	if err != nil {
		panic(err)
	}
	Database.EnsureTable("entity", &Entity{})
}

func Close() {
	Database.Close()
}
