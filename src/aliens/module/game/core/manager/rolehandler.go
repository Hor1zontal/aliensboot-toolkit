/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/7/26
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package manager

import (
	"aliens/protocol"
	"reflect"
)


func newRoleHandler(info *protocol.RoleInfo) *RoleHandler {
	handler := &RoleHandler{data : info}
	handler.Init()
	return handler
}

//角色数据管理
type RoleHandler struct {
	data *protocol.RoleInfo //
}

func (this *RoleHandler) IsRole(roleID int64) bool {
	return this.data.Id == roleID
}

func (this *RoleHandler) GetData() *protocol.RoleInfo {
	return this.data
}

//初始化
func (this *RoleHandler) Init() {
	mutable := reflect.ValueOf(this).Elem()
	params := make([]reflect.Value, 1)
	//数据管理类操作副本数据，这样更新的时候能够做比对增量更新
	params[0] = reflect.ValueOf(this.data)
	for i := 0; i < mutable.NumField(); i++ {
		f := mutable.Field(i)
		initMethod := f.Addr().MethodByName("Init")
		if initMethod.IsValid() {
			initMethod.Call(params)
		}
	}
}

//更新本地缓存
func (this *RoleHandler) Update() {
	mutable := reflect.ValueOf(this).Elem()
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(this.data)
	for i := 0; i < mutable.NumField(); i++ {
		f := mutable.Field(i)
		initMethod := f.Addr().MethodByName("Update")
		if initMethod.IsValid() {
			initMethod.Call(params)
		}
	}
}