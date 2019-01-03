package db

import (
	"github.com/KylinHe/aliensboot/database/mongo"
	"e.coding.net/aliens/aliensboot_testserver/module/passport/conf"
	"e.coding.net/aliens/aliensboot_testserver/protocol"
)

var Database *mongo.Database = &mongo.Database{}
var DatabaseHandler = Database

func Init() {
	err := Database.Init(conf.Config.Database)
	if err != nil {
		panic(err)
	}
	DatabaseHandler.EnsureTable("passport", &protocol.User{})

	//DatabaseHandler.Insert(&passport.User{Id:DatabaseHandler.GenTimestampId(&passport.User{}),Username:"hejialin",RegTime:time.Now()})
}

func Close() {
	Database.Close()
}
