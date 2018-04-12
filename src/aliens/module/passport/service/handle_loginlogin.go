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
func handleLoginLogin(request *passport.LoginLogin, result *passport.LoginLoginRet) {
	username := request.GetUsername()
	password := request.GetPassword()

	userCache := cache.GetUser(username)
	if userCache == nil {
		//TODO 后续可以做成缓存读不到去数据库并写回到缓存,要考虑数据穿透的情况
		result.Result = passport.LoginLoginRet_invalidUser
		return
	}

	passwordHash := PasswordHash(username, password)
	//密码不对
	if passwordHash != userCache.Password {
		result.Result = passport.LoginLoginRet_invalidPwd
		return
	}

	//更新ip
	//qdoc := bson.M{"_id": userCache.ID}
	//udoc := bson.M{"$set": bson.M{"ip": getNetworkAddress(network)}}
	//db.DatabaseHandler.Update(userCache.Name(), qdoc, udoc)
	result.Uid = userCache.ID
	token := NewToken()
	cache.PassportCache.SetUserToken(userCache.ID, token)
	result.Token = token
	result.Result = passport.LoginLoginRet_loginSuccess
}