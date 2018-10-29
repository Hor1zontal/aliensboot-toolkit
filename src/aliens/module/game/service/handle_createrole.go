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
	"aliens/module/game/core"
	"aliens/protocol"
)

//
func handleCreateRole(authID int64, gateID string, request *protocol.CreateRole, response *protocol.CreateRoleRet) {
	userSession := core.UserManager.EnsureUser(authID)
	role := userSession.CreateRole(request.GetRole())
	response.Role = role
}
