/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/3/29
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package handler

import "aliens/protocol/passport"

type PassportLoginServiceProxy struct {

}

//登录账号服务器请求
type PassportLoginService struct {
}

func (service *PassportLoginService) Handle(request *passport.LoginLogin, response *passport.LoginLoginRet) error {

	request.GetLoginLogin
	//message := request.GetLoginLogin()
	//username := message.GetUsername()
	//passwd := message.GetPassword()
	//result := &protocol.LoginLoginRet{}
	//response.LoginLoginRet = result
	//userCache := cache.GetUser(username)
	//if userCache == nil {
	//	//TODO 后续可以做成缓存读不到去数据库并写回到缓存,要考虑数据穿透的情况
	//	result.Result = Enum(protocol.Login_Result_invalidUser)
	//	return
	//}
	//passwordHash := helper.PasswordHash(username, passwd)
	////密码不对
	//if passwordHash != userCache.Password {
	//	result.Result = Enum(protocol.Login_Result_invalidPwd)
	//	return
	//}
	//gameServer := helper.AllocGameServer(userCache.ID)
	//if gameServer == "" {
	//	result.Result = Enum(protocol.Login_Result_invalidGameServer)
	//	return
	//}
	//
	////更新ip
	//qdoc := bson.M{"_id": userCache.ID}
	//udoc := bson.M{"$set": bson.M{"ip": getNetworkAddress(network)}}
	//db.DatabaseHandler.Update(userCache.Name(), qdoc, udoc)
	//
	//result.Uid = proto.Int32(userCache.ID)
	//token := util.Rand().Hex()
	//cache.UserCache.SetUserToken(userCache.ID, token)
	//result.Token = proto.String(token)
	//result.GameServer = proto.String(gameServer)
	//result.Result = Enum(protocol.Login_Result_loginSuccess)
	return nil
}
