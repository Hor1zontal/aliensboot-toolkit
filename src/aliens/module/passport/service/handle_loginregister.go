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
func handleLoginRegister(request *passport.LoginRegister, result *passport.LoginRegisterRet) {
	username := request.GetUsername()
	passwd := request.GetPassword()
	if cache.PassportCache.IsUsernameExist(username) {
		result.Msg = "用户名已存在"
		result.Result = passport.LoginRegisterRet_userExists
		return
	}

	passwd = PasswordHash(username, passwd)
	//TODO 有风险最好查询 数据库再加一层判断
	userCache := cache.NewUser(username, passwd, "ip address", "", "", "", "")

	result.Result = passport.LoginRegisterRet_registerSuccess
	result.Uid = userCache.Id
	token := NewToken()
	cache.PassportCache.SetUserToken(userCache.Id, token)
	result.Token = token
}
