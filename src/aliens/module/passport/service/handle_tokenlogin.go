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
	"aliens/protocol/passport"
	"aliens/module/passport/cache"
)


//
func handleTokenLogin(request *passport.TokenLogin, response *passport.TokenLoginRet) int64 {
	if cache.PassportCache.GetUserToken(request.GetUid()) != request.GetToken() {
		response.Result = passport.LoginResult_tokenExpire
		return 0
	}
	response.Result = passport.LoginResult_loginSuccess
	return request.GetUid()
}
