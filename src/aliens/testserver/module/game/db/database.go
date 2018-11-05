package db

import (
	"aliens/aliensbot/database/mongo"
	"aliens/testserver/module/game/conf"
	"aliens/testserver/protocol"
)

var Database *mongo.Database = &mongo.Database{}

func Init() {
	err := Database.Init(conf.Config.Database)
	if err != nil {
		panic(err)
	}
	Database.EnsureTable("game_user", &protocol.GameUser{})
	Database.EnsureTable("role", &protocol.RoleInfo{})

}

func Close() {
	Database.Close()
}
