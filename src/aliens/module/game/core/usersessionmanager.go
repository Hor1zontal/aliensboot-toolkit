/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/7/25
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core

import (
	"aliens/module/game/core/manager"
	"aliens/exception"
	"aliens/protocol"
)

var UserManager = &userManager{users:make(map[int64]*UserSession)}

type userManager struct {
	users map[int64]*UserSession
}


//加载用户基础数据
func (this *userManager) EnsureUser(uid int64) *UserSession {
	session := this.users[uid]
	if session == nil {
		session = newUserSession(uid)
		this.users[uid] = session
	}
	return session
}

//获取角色数据处理句柄

func (this *userManager) EnsureRoleHandler(uid int64) *manager.RoleHandler {
	session := this.EnsureUser(uid)
	handler := session.GetActiveRoleHandler()
	if handler == nil {
		exception.GameException(protocol.Code_RoleNotSelect)
	}
	return handler
}
