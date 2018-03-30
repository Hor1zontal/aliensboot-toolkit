package service

import (
	"aliens/protocol"
	"aliens/cluster/message"
	"aliens/module/cluster"
)

var PassportService = message.NewLocalService(cluster.SERVICE_PASSPORT)

func init() {
	PassportService.RegisterHandler(6, new(PassportRegisterService)) //注册
	PassportService.RegisterHandler(7, new(PassportLoginService))    //登录
}


//登录账号服务器请求
type PassportLoginService struct {
}

func (service *PassportLoginService) Request(request interface{}, response interface{}) error {



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

func Enum(v protocol.Login_Result) *protocol.Login_Result {
	return &v
}

func RegEnum(v protocol.Register_Result) *protocol.Register_Result {
	return &v
}

//账号服务器请求
type PassportRegisterService struct {
}

func (service *PassportRegisterService) Request(request interface{}, response interface{}) error {
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
	return nil
}
