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
	"aliens/exception"
	"aliens/log"
	"aliens/module/game/db"
	"aliens/protocol"
)

func NewUserDataManager(uid int64) *UserDataManager {
	dataManager := &UserDataManager{}
	var user = &protocol.GameUser{Uid: uid}
	err := db.Database.QueryOne(user)
	if err != nil {
		//创建数据
	}
	dataManager.user = user
	return dataManager
}

//角色管理容器
type UserDataManager struct {
	user       *protocol.GameUser //用户游戏信息 拥有的角色
	activeRole *RoleHandler       //当前的角色处理句柄
}

func (this *UserDataManager) GetActiveRoleHandler() *RoleHandler {
	return this.activeRole
}

func (this *UserDataManager) LoginRole(roleID int64) *RoleHandler {
	if !this.HaveRole(roleID) {
		exception.GameException(protocol.Code_RoleNotExists)
	}
	//加载当前玩家缓存
	if this.activeRole == nil || !this.activeRole.IsRole(roleID) {
		roleInfo := &protocol.RoleInfo{Id: roleID}
		err := db.Database.QueryOne(roleInfo)
		if err != nil {
			log.Debugf("query role exception %v", err)
			exception.GameException(protocol.Code_DBExcetpion)
		}
		this.activeRole = newRoleHandler(roleInfo)
	}
	return this.activeRole
}

func (this *UserDataManager) HaveRole(roleID int64) bool {
	for _, role := range this.user.Roles {
		if role.Id == roleID {
			return true
		}
	}
	return false
}

func (this *UserDataManager) GetUserData() *protocol.GameUser {
	return this.user
}

func (this *UserDataManager) CreateRole(role *protocol.Role) *protocol.Role {
	roleInfo := &protocol.RoleInfo{}
	err := db.Database.Insert(roleInfo)
	if err != nil {
		exception.GameException(protocol.Code_DBExcetpion)
	}
	role.Id = roleInfo.Id
	this.user.Roles = append(this.user.Roles, role)
	this.activeRole = newRoleHandler(roleInfo)
	return role
}
