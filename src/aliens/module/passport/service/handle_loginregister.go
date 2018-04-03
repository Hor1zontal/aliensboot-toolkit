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

import "aliens/protocol/passport"


//
func handleLoginRegister(request *passport.LoginRegister, response *passport.LoginRegisterRet) {
	//message := request.GetLoginRegister()
	//username := message.GetUsername()
	//passwd := message.GetPassword()
	//result := &protocol.LoginRegisterRet{}
	//response.LoginRegisterRet = result
	//if cache.UserCache.IsUsernameExist(username) {
	//	result.Msg = proto.String("用户名已存在")
	//	result.Result = RegEnum(protocol.Register_Result_userExists)
	//	return
	//}
	//
	//passwd = helper.PasswordHash(username, passwd)
	////TODO 有风险最好数据库再加一层判断
	//userCache := cache.NewUser(username, passwd, getNetworkAddress(network), "", "", "", "")
	//gameServer := helper.AllocGameServer(0)
	//if gameServer == "" {
	//	result.Result = RegEnum(protocol.Register_Result_invalidServer)
	//	return
	//}
	//result.Result = RegEnum(protocol.Register_Result_registerSuccess)
	//result.Uid = proto.Int32(userCache.ID)
	//token := util.Rand().Hex()
	//cache.UserCache.SetUserToken(userCache.ID, token)
	//result.Token = proto.String(token)
	//result.GameServer = proto.String(gameServer)
}
