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
	"aliens/module/passport/cache"
	"aliens/protocol"
)



func handleTokenLogin(request *protocol.TokenLogin, response *protocol.TokenLoginRet) int64 {
	token, _ := cache.PassportCache.GetUserToken(request.GetUid())
	if token != request.GetToken() {
		response.Result = protocol.LoginResult_tokenExpire
		return 0
	}
	response.Result = protocol.LoginResult_loginSuccess
	request.Uid = 2
	return request.GetUid()
}
