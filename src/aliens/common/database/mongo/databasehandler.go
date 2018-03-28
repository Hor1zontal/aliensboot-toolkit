package mongo

import (
	"aliens/common/database"
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	ID_STORE string = "_id"
)

//新增自增长键
func (this *Database) EnsureCounter(data database.IData) {
	this.validateConnection()
	this.dbContext.EnsureCounter(this.dbName, ID_STORE, data.Name())
}

//确保索引
func (this *Database) EnsureUniqueIndex(data database.IData, key string) {
	this.validateConnection()
	this.dbContext.EnsureUniqueIndex(this.dbName, data.Name(), []string{key})
}

func (this *Database) EnsureTable(data database.IData) {
	this.validateConnection()
	dataType := reflect.TypeOf(data).Elem()
	field, ok := dataType.FieldByName("id")
	if ok {
		value := field.Tag.Get("gorm")
		if strings.Contains(value, "AUTO_INCREMENT") {
			this.dbContext.EnsureCounter(this.dbName, ID_STORE, data.Name())
		}
	}
}

func (this *Database) Related(data interface{}, relateData interface{}, relateTableName string, relateKey string) error {
	//mongo采用树形结构，不用建立关系
	return nil
}

//自增长id
func (this *Database) GenId(data database.IData) int32 {
	this.validateConnection()
	newid, _ := this.dbContext.NextSeq(this.dbName, ID_STORE, data.Name())
	return int32(newid)
}

func (this *Database) GenTimestampId(data database.IData) int64 {
	this.validateConnection()
	newid, _ := this.dbContext.NextSeq(this.dbName, ID_STORE, data.Name())
	idStr := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(newid)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		panic("invalid timestamp id:" + idStr)
	}

	return id
}

//插入新数据
func (this *Database) Insert(data database.IData) error {
	this.validateConnection()
	return this.database.C(data.Name()).Insert(data)
}

//查询所有数据
func (this *Database) QueryAll(data database.IData, result interface{}) error {
	this.validateConnection()
	err := this.database.C(data.Name()).Find(nil).All(result)
	if err != nil {
		log.Error("%v", err)
	}
	return err
}

//查询单条记录
func (this *Database) QueryOne(data database.IData) error {
	this.validateConnection()
	err := this.database.C(data.Name()).FindId(data.GetID()).One(data)
	if err != nil {
		log.Error("%v", err)
	}
	return err
}

func (this *Database) DeleteOne(data database.IData) error {
	this.validateConnection()
	err := this.database.C(data.Name()).RemoveId(data.GetID())
	if err != nil {
		log.Error("%v", err)
	}
	return err
}

func (this *Database) DeleteOneCondition(data database.IData, selector interface{}) error {
	this.validateConnection()
	err := this.database.C(data.Name()).Remove(selector)
	if err != nil {
		log.Error("%v", err)
	}
	return err
}

//查询单条记录
func (this *Database) IDExist(data database.IData) bool {
	this.validateConnection()
	count, err := this.database.C(data.Name()).FindId(data.GetID()).Count()
	if err != nil {
		log.Error("%v", err)
		return false
	}
	return count != 0
}

//按条件多条查询
func (this *Database) QueryAllCondition(data database.IData, condition string, value interface{}, result interface{}) error {
	err := this.database.C(data.Name()).Find(bson.M{condition: value}).All(result)
	if err != nil {
		log.Error("%v", err)
	}
	return err
}

//按条件单条查询
func (this *Database) QueryOneCondition(data database.IData, condition string, value interface{}) error {
	this.validateConnection()
	err := this.database.C(data.Name()).Find(bson.M{condition: value}).One(data)
	if err != nil {
		log.Error("%v", err)
	}
	return err
}

//更新单条数据
func (this *Database) UpdateOne(data database.IData) error {
	this.validateConnection()
	err := this.database.C(data.Name()).UpdateId(data.GetID(), data)
	if err != nil {
		log.Error("%v", err)
	}
	return err
}

func (this *Database) ForceUpdateOne(data database.IData) error {
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
