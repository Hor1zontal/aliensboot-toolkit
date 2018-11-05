package mongo

import (
	"gopkg.in/mgo.v2/bson"
	"reflect"
	//"strconv"
	"strings"
	//"time"
	//"aliens/aliensbot/log"
	"aliens/aliensbot/database/dbconfig"
	"github.com/pkg/errors"
)

const (
	ID_STORE      string = "_id"
	ID_FIELD_NAME string = "ID"
)

//获取表格名和id值
func (this *Database) GetTableMeta(data interface{}) (*dbconfig.TableMeta, error) {
	tableType := reflect.TypeOf(data)
	result, ok := this.tableMetas[tableType]
	if !ok {
		return result, errors.New("un expect db collection " + tableType.String())
	}
	return result, nil
}

//func (this *Database) GetID(data interface{}) interface{} {
//	tableMeta, err := this.GetTableMeta(data)
//	if err != nil {
//		return -1
//	}
//	return this.reflectID(data, tableMeta.IDName)
//}

func (this *Database) reflectID(data interface{}, idName string) interface{} {
	return reflect.ValueOf(data).Elem().FieldByName(idName).Interface()
}

//新增自增长键
//func (this *Database) EnsureCounter(data interface{}) {
//	this.validateConnection()
//	tableMeta := this.GetTableMeta(data)
//	this.dbContext.EnsureCounter(this.dbName, ID_STORE, tableMeta)
//}

//确保索引
//func (this *Database) EnsureUniqueIndex(data interface{}, key string) error {
//	this.validateConnection()
//	tableMeta, err := this.GetTableMeta(data)
//	if err != nil {
//		return err
//	}
//	return this.dbContext.EnsureUniqueIndex(this.dbName, tableMeta, []string{key})
//}

func (this *Database) EnsureTable(name string, data interface{}) error {
	this.validateConnection()
	tableType := reflect.TypeOf(data)
	if tableType == nil || tableType.Kind() != reflect.Ptr {
		return errors.New("table data pointer required")
	}

	meta := &dbconfig.TableMeta{Name: name}
	dataType := tableType.Elem()
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		uniqueValue := field.Tag.Get("unique")
		if strings.Contains(uniqueValue, "true") {
			key := field.Tag.Get("bson")
			if key != "" {
				this.dbContext.EnsureUniqueIndex(this.dbName, name, []string{key})
			}
		} else {
			idKey := field.Tag.Get("bson")
			if idKey == ID_STORE {
				meta.IDName = field.Name
				value := field.Tag.Get("gorm")
				if strings.Contains(value, "AUTO_INCREMENT") {
					meta.AutoIncrement = true
					this.dbContext.EnsureCounter(this.dbName, ID_STORE, name)
				}
			}
		}
	}
	if meta.IDName == "" {
		return errors.New("bson:_id is not found in " + name + " tag")
	}
	this.tableMetas[tableType] = meta
	return nil
}

func (this *Database) Related(data interface{}, relateData interface{}, relateTableName string, relateKey string) error {
	//mongo采用树形结构，不用建立关系
	return nil
}

//自增长id
//func (this *Database) GenId(data interface{}) (int32, error) {
//	this.validateConnection()
//	tableMeta, err := this.GetTableMeta(data)
//	if err != nil {
//		return -1, err
//	}
//	newid, _ := this.dbContext.NextSeq(this.dbName, ID_STORE, tableMeta)
//	return int32(newid), nil
//}

//func (this *Database) GenTimestampId(data interface{}) (int64, error) {
//	this.validateConnection()
//	tableMeta, err := this.GetTableMeta(data)
//	if err != nil {
//		return -1, err
//	}
//	newid, _ := this.dbContext.NextSeq(this.dbName, ID_STORE, tableMeta)
//	idStr := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(newid)
//	return strconv.ParseInt(idStr, 10, 64)
//}

//插入新数据
func (this *Database) Insert(data interface{}) error {
	this.validateConnection()
	tableMeta, err := this.GetTableMeta(data)
	if err != nil {
		return err
	}

	if tableMeta.AutoIncrement {
		newid, err1 := this.dbContext.NextSeq(this.dbName, ID_STORE, tableMeta.Name)
		if err1 != nil {
			return err1
		}
		reflect.ValueOf(data).Elem().FieldByName(tableMeta.IDName).SetInt(int64(newid))
	}
	return this.database.C(tableMeta.Name).Insert(data)
}

func (this *Database) QueryAllLimit(data interface{}, result interface{}, limit int, callback func(interface{}) bool) error {
	this.validateConnection()
	tableMeta, err := this.GetTableMeta(data)
	if err != nil {
		return err
	}
	skip := 0
	for {
		err := this.database.C(tableMeta.Name).Find(nil).Limit(limit).Skip(skip).All(result)
		if err != nil {
			return err
		}
		skip += limit
		if callback(result) {
			return nil
		}
	}
}

//按条件多条查询
func (this *Database) QueryAllConditionLimit(data interface{}, condition string, value interface{}, result interface{}, limit int, callback func(interface{}) bool) error {
	this.validateConnection()
	tableMeta, err := this.GetTableMeta(data)
	if err != nil {
		return err
	}
	skip := 0
	for {
		err := this.database.C(tableMeta.Name).Find(bson.M{condition: value}).Limit(limit).Skip(skip).All(result)
		if err != nil {
			return err
		}
		skip += limit
		if callback(result) {
			return nil
		}
	}
}

//查询所有数据
func (this *Database) QueryAll(data interface{}, result interface{}) error {
	this.validateConnection()
	tableMeta, err := this.GetTableMeta(data)
	if err != nil {
		return err
	}
	return this.database.C(tableMeta.Name).Find(nil).All(result)
}

//查询单条记录
func (this *Database) QueryOne(data interface{}) error {
	this.validateConnection()
	tableMeta, err := this.GetTableMeta(data)
	if err != nil {
		return err
	}
	return this.database.C(tableMeta.Name).FindId(this.reflectID(data, tableMeta.IDName)).One(data)
}

func (this *Database) DeleteOne(data interface{}) error {
	this.validateConnection()
	tableMeta, err := this.GetTableMeta(data)
	if err != nil {
		return err
	}
	return this.database.C(tableMeta.Name).RemoveId(this.reflectID(data, tableMeta.IDName))
}

func (this *Database) DeleteOneCondition(data interface{}, selector interface{}) error {
	this.validateConnection()
	tableMeta, err := this.GetTableMeta(data)
	if err != nil {
		return err
	}
	return this.database.C(tableMeta.Name).Remove(selector)
}

//查询单条记录
func (this *Database) IDExist(data interface{}) (bool, error) {
	this.validateConnection()
	tableMeta, err := this.GetTableMeta(data)
	if err != nil {
		return false, err
	}
	count, err1 := this.database.C(tableMeta.Name).FindId(this.reflectID(data, tableMeta.IDName)).Count()
	return count != 0, err1
}

//按条件多条查询
func (this *Database) QueryAllCondition(data interface{}, condition string, value interface{}, result interface{}) error {
	this.validateConnection()
	tableMeta, err := this.GetTableMeta(data)
	if err != nil {
		return err
	}
	return this.database.C(tableMeta.Name).Find(bson.M{condition: value}).All(result)
}

//按条件单条查询
func (this *Database) QueryOneCondition(data interface{}, condition string, value interface{}) error {
	this.validateConnection()
	tableMeta, err := this.GetTableMeta(data)
	if err != nil {
		return err
	}
	return this.database.C(tableMeta.Name).Find(bson.M{condition: value}).One(data)
}

//更新单条数据
func (this *Database) UpdateOne(data interface{}) error {
	this.validateConnection()
	tableMeta, err := this.GetTableMeta(data)
	if err != nil {
		return err
	}
	return this.database.C(tableMeta.Name).UpdateId(this.reflectID(data, tableMeta.IDName), data)
}

func (this *Database) ForceUpdateOne(data interface{}) error {
	result, err := this.IDExist(data)
	if err != nil {
		return err
	}
	if result {
		return this.UpdateOne(data)
	} else {
		return this.Insert(data)
	}
}

//原生的更新语句
//TODO 需要拓展到内存映射修改，减少开发量
func (this *Database) Update(collection string, selector interface{}, update interface{}) error {
	this.validateConnection()
	return this.database.C(collection).Update(selector, update)
}
