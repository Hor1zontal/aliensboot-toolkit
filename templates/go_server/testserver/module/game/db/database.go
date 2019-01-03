package db

import (
	"github.com/KylinHe/aliensboot/database/mongo"
	"e.coding.net/aliens/aliensboot_testserver/module/game/conf"
	"e.coding.net/aliens/aliensboot_testserver/protocol"
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
