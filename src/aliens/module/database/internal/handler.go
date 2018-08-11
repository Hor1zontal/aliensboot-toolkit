package internal

import (
	"aliens/database"
	"aliens/log"
	"time"
	"reflect"
	"aliens/module/database/conf"
	"aliens/module/database/constant"
)

func init() {
	// 向当前模块注册客户端发送的消息处理函数 handleMessage
	skeleton.RegisterChanRPC(constant.DB_COMMAND_INSERT, handleInsert)
	skeleton.RegisterChanRPC(constant.DB_COMMAND_UPDATE, handleUpdate)
	skeleton.RegisterChanRPC(constant.DB_COMMAND_DELETE, handleDelete)
	skeleton.RegisterChanRPC(constant.DB_COMMAND_FUPDATE, forceUpdate)
	skeleton.RegisterChanRPC(constant.DB_COMMAND_CONDITION_UPDATE, conditionUpdate)
}

func handleDelete(args []interface{}) {
	handler := args[1].(database.IDatabaseHandler)
	starTime := time.Now()
	err := handler.DeleteOne(args[0])
	debugLog("delete", args[0], starTime, err)
}

func handleInsert(args []interface{}) {
	handler := args[1].(database.IDatabaseHandler)
	starTime := time.Now()
	err := handler.Insert(args[0])
	debugLog("insert", args[0], starTime, err)
}

func handleUpdate(args []interface{}) {
	handler := args[1].(database.IDatabaseHandler)
	starTime := time.Now()
	err := handler.UpdateOne(args[0])
	debugLog("update", args[0], starTime, err)
}

func forceUpdate(args []interface{}) {
	handler := args[1].(database.IDatabaseHandler)
	starTime := time.Now()
	err := handler.ForceUpdateOne(args[0])
	debugLog("force update", args[0], starTime, err)
}

func conditionUpdate(args []interface{}) {
	handler := args[3].(database.IDatabaseHandler)
	starTime := time.Now()
	err := handler.Update(args[0].(string), args[1], args[2])
	debugLog("condition update", args[0], starTime, err)
}

func debugLog(opt string, data interface{}, startTime time.Time, err error) {
	if conf.DB_DEBUG {
		typeName := reflect.TypeOf(data).Name()
		if err != nil {
			log.Debug("[%v] %v err %v", typeName, err)
		}
		duration := time.Now().Sub(startTime)
		if duration.Seconds() >= conf.DB_SIGAL_TIMEOUT {
			log.Debug("[%v] %v too long %v",opt, typeName, duration.Seconds())
		}
	}
}




