package mongo

import (
	"aliens/database"
	"github.com/name5566/leaf/db/mongodb"
	"gopkg.in/mgo.v2"
	"aliens/database/dbconfig"
	"reflect"
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

	table     map[reflect.Type]string
	tableIDName  map[reflect.Type]string //表格id字段名
}

//初始化连接数据库
func (this *Database) Init(config dbconfig.DBConfig) error {
	this.dbName = config.Name
	if config.MaxSession == 0 {
		config.MaxSession = 100
	}
	c, err := mongodb.Dial(config.Address, int(config.MaxSession))
	if err != nil {
		return err
	}
	this.table = make(map[reflect.Type]string)
	this.tableIDName = make(map[reflect.Type]string)
	this.dbContext = c
	this.dbSession = this.dbContext.Ref()
	this.database = this.dbSession.DB(config.Name)
	//if (this.auth != nil) {
	//	return this.database.Login(this.auth.Username, this.auth.Password)
	//}
	return nil
}

//初始化账号密码信息
//func (this *Database) auth(username string, password string) {
//	if username != "" {
//		this.auth = &database.Authority{username, password}
//	}
//}

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
