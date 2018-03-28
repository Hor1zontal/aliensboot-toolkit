package mongo

import (
	"aliens/common/database"
	"github.com/name5566/leaf/db/mongodb"
	"gopkg.in/mgo.v2"
)

//type DatabaseFactory struct {
//
//}

//func (this DatabaseFactory) Create() database.IDatabase {
//	//TODO 根据参数定制
//	return &Database{}
//}

type Database struct {
	dbName    string
	dbContext *mongodb.DialContext
	dbSession *mongodb.Session
	database  *mgo.Database
	auth      *database.Authority
}

//初始化连接数据库
func (this *Database) Init(host string, port int, dbName string) error {
	this.dbName = dbName
	c, err := mongodb.Dial(host, 100)
	if err != nil {
		return err
	}
	this.dbContext = c
	this.dbSession = this.dbContext.Ref()
	this.database = this.dbSession.DB(dbName)
	//if (this.auth != nil) {
	//	return this.database.Login(this.auth.Username, this.auth.Password)
	//}
	return nil
}

//初始化账号密码信息
func (this *Database) Auth(username string, password string) {
	if username != "" {
		this.auth = &database.Authority{username, password}
	}
}

func (this *Database) validateConnection() {
	if !this.isConnect() {
		panic("database is not connection")
	}
}

func (this *Database) isConnect() bool {
	return this.database != nil
}

func (this *Database) Close() {
	if this.dbContext == nil {
		return
	}
	if this.dbSession != nil {
		this.dbContext.UnRef(this.dbSession)
	}
	this.dbContext.Close()
}

func (this *Database) GetHandler() database.IDatabaseHandler {
	return this
}
