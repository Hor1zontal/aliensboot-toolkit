package mongo

import (
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strconv"
	"strings"
	"time"
	"aliens/log"
)

const (
	ID_STORE string = "_id"
	ID_FIELD_NAME string = "Id"
)


//获取表格名和id值
func (this *Database) GetTableName(data interface{}) string {
	return this.table[reflect.TypeOf(data)]
}

func (this *Database) GetID(data interface{}) interface{} {
	return reflect.ValueOf(data).FieldByName(ID_FIELD_NAME).Interface()
}

//新增自增长键
//func (this *Database) EnsureCounter(data interface{}) {
//	this.validateConnection()
//	tableName := this.GetTableName(data)
//	this.dbContext.EnsureCounter(this.dbName, ID_STORE, tableName)
//}

//确保索引
func (this *Database) EnsureUniqueIndex(data interface{}, key string) {
	this.validateConnection()
	tableName := this.GetTableName(data)
	this.dbContext.EnsureUniqueIndex(this.dbName, tableName, []string{key})
}

func (this *Database) EnsureTable(name string, data interface{}) {
	this.validateConnection()
	tableType := reflect.TypeOf(data)
	if tableType == nil || tableType.Kind() != reflect.Ptr {
		log.Fatal("table data pointer required")
	}
	this.table[tableType] = name

	dataType := tableType.Elem()
	field, ok := dataType.FieldByName(ID_FIELD_NAME)
	if ok {
		value := field.Tag.Get("gorm")
		if strings.Contains(value, "AUTO_INCREMENT") {
			this.dbContext.EnsureCounter(this.dbName, ID_STORE, name)
		}
	}
}

func (this *Database) Related(data interface{}, relateData interface{}, relateTableName string, relateKey string) error {
	//mongo采用树形结构，不用建立关系
	return nil
}

//自增长id
func (this *Database) GenId(data interface{}) int32 {
	this.validateConnection()
	tableName := this.GetTableName(data)
	newid, _ := this.dbContext.NextSeq(this.dbName, ID_STORE, tableName)
	return int32(newid)
}

func (this *Database) GenTimestampId(data interface{}) int64 {
	this.validateConnection()
	tableName := this.GetTableName(data)
	newid, _ := this.dbContext.NextSeq(this.dbName, ID_STORE, tableName)
	idStr := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(newid)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		panic("invalid timestamp id:" + idStr)
	}

	return id
}

//插入新数据
func (this *Database) Insert(data interface{}) error {
	this.validateConnection()
	tableName := this.GetTableName(data)
	log.Debug(tableName)
	return this.database.C(tableName).Insert(data)
}

//查询所有数据
func (this *Database) QueryAll(data interface{}, result interface{}) error {
	this.validateConnection()
	tableName := this.GetTableName(data)
	err := this.database.C(tableName).Find(nil).All(result)
	if err != nil {
		log.Error("%v", err)
	}
	return err
}

//查询单条记录
func (this *Database) QueryOne(data interface{}) error {
	this.validateConnection()
	tableName := this.GetTableName(data)
	err := this.database.C(tableName).FindId(this.GetID(data)).One(data)
	if err != nil {
		log.Error("%v", err)
	}
	return err
}

func (this *Database) DeleteOne(data interface{}) error {
	this.validateConnection()
	tableName := this.GetTableName(data)
	err := this.database.C(tableName).RemoveId(this.GetID(data))
	if err != nil {
		log.Error("%v", err)
	}
	return err
}

func (this *Database) DeleteOneCondition(data interface{}, selector interface{}) error {
	this.validateConnection()
	tableName := this.GetTableName(data)
	err := this.database.C(tableName).Remove(selector)
	if err != nil {
		log.Error("%v", err)
	}
	return err
}

//查询单条记录
func (this *Database) IDExist(data interface{}) bool {
	this.validateConnection()
	tableName := this.GetTableName(data)
	count, err := this.database.C(tableName).FindId(this.GetID(data)).Count()
	if err != nil {
		log.Error("%v", err)
		return false
	}
	return count != 0
}

//按条件多条查询
func (this *Database) QueryAllCondition(data interface{}, condition string, value interface{}, result interface{}) error {
	this.validateConnection()
	tableName := this.GetTableName(data)
	err := this.database.C(tableName).Find(bson.M{condition: value}).All(result)
	if err != nil {
		log.Error("%v", err)
	}
	return err
}

//按条件单条查询
func (this *Database) QueryOneCondition(data interface{}, condition string, value interface{}) error {
	this.validateConnection()
	tableName := this.GetTableName(data)
	err := this.database.C(tableName).Find(bson.M{condition: value}).One(data)
	if err != nil {
		log.Error("%v", err)
	}
	return err
}

//更新单条数据
func (this *Database) UpdateOne(data interface{}) error {
	this.validateConnection()
	tableName := this.GetTableName(data)
	err := this.database.C(tableName).UpdateId(this.GetID(data), data)
	if err != nil {
		log.Error("%v", err)
	}
	return err
}

func (this *Database) ForceUpdateOne(data interface{}) error {
	if this.IDExist(data) {
		return this.UpdateOne(data)
	} else {
		return this.Insert(data)
	}
}

//原生的更新语句
//TODO 需要拓展到内存映射修改，减少开发量
func (this *Database) Update(collection string, selector interface{}, update interface{}) error {
	this.validateConnection()
	err := this.database.C(collection).Update(selector, update)
	if err != nil {
		log.Error("%v", err)
	}
	return err
}
