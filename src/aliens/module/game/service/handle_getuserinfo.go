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
	"aliens/module/game/core"
	"aliens/protocol"
)


//
func handleGetUserInfo(authID int64, request *protocol.GetUserInfo, response *protocol.GetUserInfoRet) {
	userSession := core.UserManager.EnsureUser(authID)
	response.User = userSession.GetUserData()
}
