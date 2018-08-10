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


//
func handleLoginLogin(request *protocol.LoginLogin, result *protocol.LoginLoginRet) int64 {
	username := request.GetUsername()
	password := request.GetPassword()

	userCache := cache.GetUser(username)
	if userCache == nil {
		passwordHash := PasswordHash(username, password)


		userCache = cache.NewUser(username, passwordHash, "ip address", "", "", "", "")
		//result.Result = passport.LoginLoginRet_invalidUser
		//return
	}

	passwordHash := PasswordHash(username, password)
	//密码不对
	if passwordHash != userCache.Password {
		result.Result = protocol.LoginResult_invalidPwd
		return 0
	}

	//更新ip
	//qdoc := bson.M{"_id": userCache.ID}
	//udoc := bson.M{"$set": bson.M{"ip": getNetworkAddress(network)}}
	//db.DatabaseHandler.Update(userCache.Name(), qdoc, udoc)
	result.Uid = userCache.GetId()
	token := NewToken()
	cache.PassportCache.SetUserToken(userCache.GetId(), token)
	result.Token = token
	result.Result = protocol.LoginResult_loginSuccess
	return result.GetUid()
}