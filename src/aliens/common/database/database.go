package database

import "aliens/common/database/dbconfig"

//数据库抽象层 适配其他数据库
type IDatabase interface {
	Init(config dbconfig.DBConfig) error //初始化数据库
	//Auth(username string, password string)           //登录信息
	Close()                                          //关闭数据库
	GetHandler() IDatabaseHandler                    //获取数据库处理类
}

type IDatabaseFactory interface {
	create() IDatabase
}

type Authority struct {
	Username string
	Password string
}

//数据库handler
type IDatabaseHandler interface {
	EnsureTable(data IData)                                                                           //确保表存在
	EnsureUniqueIndex(data IData, name string)                                                        //确保索引
	Related(data interface{}, relateData interface{}, relateTableName string, relateKey string) error //创建依赖关系
	GenId(data IData) int32
	GenTimestampId(data IData) int64
	Insert(data IData) error
	QueryAll(data IData, result interface{}) error
	QueryAllCondition(data IData, condition string, value interface{}, result interface{}) error
	IDExist(data IData) bool
	QueryOne(data IData) error
	QueryOneCondition(data IData, condition string, value interface{}) error
	DeleteOne(data IData) error
	DeleteOneCondition(data IData, selector interface{}) error
	UpdateOne(data IData) error
	ForceUpdateOne(data IData) error //强制更新。不存在就插入
	Update(collection string, selector interface{}, update interface{}) error
}

type IData interface {
	Name() string       //数据名称
	GetID() interface{} //获取数据索引ID
}

type IRelatedData interface {
	RelateLoad(handler IDatabaseHandler) //注入关联数据
	RelateSave(handler IDatabaseHandler) //保存关联数据
}
