package database

import "aliens/database/dbconfig"

//数据库抽象层 适配其他数据库
type IDatabase interface {
	Init(config dbconfig.DBConfig) error //初始化数据库
	//auth(username string, password string)           //登录信息
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
	GetTableMeta(data interface{}) (*dbconfig.TableMeta, error)
	//GetTableName(data interface{}) (string, error)
	EnsureTable(name string, data interface{}) error //确保表存在
	//GetID(data interface{}) interface{}
	//EnsureUniqueIndex(data interface{}, name string) error                                                       //确保索引
	Related(data interface{}, relateData interface{}, relateTableName string, relateKey string) error //创建依赖关系
	//GenId(data interface{}) (int32, error)
	//GenTimestampId(data interface{}) (int64, error)
	Insert(data interface{}) error
	QueryAll(data interface{}, result interface{}) error
	QueryAllCondition(data interface{}, condition string, value interface{}, result interface{}) error
	QueryAllLimit(data interface{}, result interface{}, limit int, callback func(interface{}) bool) error
	QueryAllConditionLimit(data interface{}, condition string, value interface{}, result interface{}, limit int, callback func(interface{}) bool) error
	QueryOne(data interface{}) error
	QueryOneCondition(data interface{}, condition string, value interface{}) error
	IDExist(data interface{}) (bool, error)
	DeleteOne(data interface{}) error
	DeleteOneCondition(data interface{}, selector interface{}) error
	UpdateOne(data interface{}) error
	ForceUpdateOne(data interface{}) error //强制更新。不存在就插入
	Update(collection string, selector interface{}, update interface{}) error
}

//type interface{} interface {
//	//Name() string       //数据名称
//	//GetID() interface{} //获取数据索引ID
//}

type IRelatedData interface {
	RelateLoad(handler IDatabaseHandler) //注入关联数据
	RelateSave(handler IDatabaseHandler) //保存关联数据
}
