package db

import (
	"aliens/aliensbot/database/mongo"
	"aliens/testserver/module/passport/conf"
	"aliens/testserver/protocol"
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