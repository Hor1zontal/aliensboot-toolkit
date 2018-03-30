package db

import (
	"aliens/module/passport/conf"
	"aliens/common/database/mongo"
	"aliens/common/database"
)

var Database database.IDatabase = &mongo.Database{}
var DatabaseHandler = Database.GetHandler()

func Init() {
	Database.Auth(conf.Config.DBUsername, conf.Config.DBPassword)
	err := Database.Init(conf.Config.DBHost, conf.Config.DBPort, conf.Config.DBName)
	if err != nil {
		panic(err)
	}
	DatabaseHandler.EnsureTable(&DBUser{})
	DatabaseHandler.EnsureUniqueIndex(&DBUser{}, "username")

}

func Close() {
	Database.Close()
}
