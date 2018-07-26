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
)


//
func handleCreateRole(authID int64, request *game.CreateRole, response *game.CreateRoleRet) {
	userSession := core.UserManager.EnsureUser(authID)
	role := userSession.CreateRole(request.GetRole())
	response.Role = role
}
