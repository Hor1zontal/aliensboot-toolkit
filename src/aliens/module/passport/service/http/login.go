package http

import (
	"net/http"
	"time"
	"gok/passportserver/cache"
	"gok/passportserver/helper"
	"gok/log"
	"gok/passportserver/conf"
	"strconv"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"gok/passportserver/db"
)

const(
	PARAM_UID = "uid"
	PARAM_CHANNEL = "channel"
	PARAM_OPENID = "openid"
	PARAM_TIME = "time"
	PARAM_APPID = "appid"
	PARAM_NICK = "nick"
	PARAM_SIGN = "sign"
	PARAM_AMOUNT = "amount"
	PARAM_ORDER = "cpOrderId"
	PARAM_MEMO = "memo"
	PARAM_AVATAR = "avatar"
	PARAM_ACCOUNTID = "accountId"
)

type ResponseResult int

const (
	ResultSuccess                   ResponseResult = iota //0 成功
	ResultInvalidParam                             = 1001 //1001 无效的参数
	ResultInvalidGameServer                        = 1002 //1002 无可用的游戏服务器
	ResultInvalidSign                              = 1003 //1003 无效的签名
)

func parametersIsset(request *http.Request) bool {
	return request.FormValue(PARAM_CHANNEL) != "" &&
		request.FormValue(PARAM_UID) != "" &&
		request.FormValue(PARAM_APPID) != ""
}


func ChannelLogin(responseWriter http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	log.Debug("login request: %v", request.Form)
	if !parametersIsset(request) {
		sendToClient(responseWriter, GetErrorResponse(ResultInvalidParam))
		return
	}

	channel := request.FormValue(PARAM_CHANNEL)
	channelUID := request.FormValue(PARAM_UID)
	appID := request.FormValue(PARAM_APPID)
	sign := request.FormValue(PARAM_SIGN)
	openID := request.FormValue(PARAM_OPENID)
	time := request.FormValue(PARAM_TIME)
	avatar := request.FormValue(PARAM_AVATAR)


	nick := request.FormValue(PARAM_NICK)
	signText := "channel=" + channel + "&appid=" + appID + "&time=" + time + "&uid=" + channelUID + conf.Server.AppKey
	signResult := helper.MD5Hash(signText)
	if (signResult != sign) {
		log.Debug("sign error signResult %v : sign : %v", signResult, sign)
		sendToClient(responseWriter, GetErrorResponse(ResultInvalidSign))
		return
	}

	channelLogin(responseWriter, channel, channelUID, openID, nick, avatar)
}



func channelLogin(responseWriter http.ResponseWriter, channel string, channelUID string, openID string, nick string, avatar string) {
	username := channel + "_" + channelUID
	userCache := cache.GetUser(username)
	if userCache == nil {
		passwd := helper.PasswordHash(username, conf.Server.DefaultChannelPWD)
		userCache = cache.NewUser(username, passwd, "", channel, channelUID, openID, avatar)
	}

	gameServer := helper.AllocGameServer(userCache.ID)
	if gameServer == "" {
		sendToClient(responseWriter, GetErrorResponse(ResultInvalidGameServer))
		return
	}
	//头像有变更,需要更新缓存和数据库
	if (avatar != "" && avatar != userCache.Avatar) {
		cache.UserCache.SetUserAvatar(userCache.ID, avatar)
		qdoc := bson.M{"_id": userCache.ID}
		udoc := bson.M{"$set": bson.M{"avatar": avatar}}
		db.DatabaseHandler.Update(db.C_USER, qdoc, udoc)
	}
	token := helper.NewToken()
	cache.UserCache.SetUserToken(userCache.ID, token)
	cache.UserCache.SetUserNickname(userCache.ID, nick)
	log.Debug("channel login uid:%v username:%v token:%v time:%v", userCache.ID, userCache.Username, token, time.Now())
	sendToClient(responseWriter, GetSuccessResponse(userCache.ID, token, gameServer))
}


type LoginResponse struct {
	Code int   `json:"rcode"`
	Uid int32 `json:"uid"`
	Token string `json:"token"`
	GameServer string `json:"gameserver"`
}


func GetErrorResponse(code ResponseResult) string {
	return string("{\"rcode\":" + strconv.Itoa(int(code)) + "}")
}

func GetSuccessResponse(uid int32, token string, gameServer string) string {
	result, _ := json.Marshal(&LoginResponse{Code: int(ResultSuccess), Uid: uid, Token: token, GameServer: gameServer})
	return string(result)
}