package db

import (
	"aliens/database/mongo"
	"aliens/protocol"
	"aliens/module/game/conf"
)

var Database *mongo.Database = &mongo.Database{}


func Init() {
	err := Database.Init(conf.Config.Database)
	if err != nil {
		panic(err)
	}
	Database.EnsureTable("game_user", &protocol.GameUser{})
	Database.EnsureTable("role", &protocol.RoleInfo{})

	//DatabaseHandler.Insert(&passport.User{Id:DatabaseHandler.GenTimestampId(&passport.User{}),Username:"hejialin",RegTime:time.Now()})
}

func Close() {
	Database.Close()
}
