package db

import (
	"aliens/aliensbot/database/mongo"
	"aliens/mmoserver/module/game/conf"
	"aliens/mmoserver/protocol"
)

var Database *mongo.Database = &mongo.Database{}

func Init() {
	err := Database.Init(conf.Config.Database)
	if err != nil {
		panic(err)
	}

	Database.EnsureTable("role", &protocol.Role{})

}

func Close() {
	Database.Close()
}
