/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"aliens/aliensbot/exception"
	"aliens/aliensbot/log"
	"aliens/testserver/dispatch/rpc"
	"aliens/testserver/module/game/core"
	"aliens/testserver/protocol"
)

//
func handleLoginRole(authID int64, gateID string, request *protocol.LoginRole, response *protocol.LoginRoleRet) {
	//需要验权通过
	if authID == 0 {
		exception.GameException(protocol.Code_ValidateException)
	}
	userSession := core.UserManager.EnsureUser(authID)

	//通知场景服务器加载数据
	err := rpc.Scene.LoginScene("", &protocol.LoginScene{
		SpaceID:"space1",
		AuthID:authID,
		GateID:gateID,
	})

	//场景服务不可用
	if err != nil {
		log.Debugf("login scene err : %v", err)
		exception.GameException(protocol.Code_InvalidService)
	}

	response.Role = userSession.GetData()
}
