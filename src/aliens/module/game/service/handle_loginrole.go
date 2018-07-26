/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"aliens/protocol/game"
	"aliens/module/game/core"
	"time"
)


//
func handleLoginRole(authID int64, request *game.LoginRole, response *game.LoginRoleRet) {
	userSession := core.UserManager.EnsureUser(authID)
	roleHandler := userSession.LoginRole(request.GetRoleID())
	response.RoleInfo = roleHandler.GetData()
	response.ServerTime = time.Now().Unix()
}
