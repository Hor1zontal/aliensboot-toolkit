package cache

import (
	//basecache "gok/cache"
	//"gok/passportserver/conf"
	//"gok/passportserver/db"
	//"time"
	//"gok/network/exception"
	//"gok/log"
	"aliens/module/passport/db"
	"time"
)

//var UserCache = basecache.NewUserCacheManager()

func Init() {
	//UserCache.Init1(conf.Server.RedisAddress, conf.Server.RedisPassword,
	//	conf.Server.RedisMaxActive, conf.Server.RedisMaxIdle, conf.Server.RedisIdleTimeout)
	////其他服务器加载了缓存就不需要再加载到缓存了
	//if UserCache.SetNX(basecache.FLAG_LOADUSER, 1) {
	//	var users []*db.DBUser
	//	db.DatabaseHandler.QueryAll(&db.DBUser{}, &users)
	//	for _, user := range users {
	//		UserCache.SetUsernameUidMapping(user.Username, user.ID)
	//		UserCache.HSetUser(user.ID, user)
	//	}
	//}
}

func Close() {

	//清除所有缓存数据
	//UserCache.Close()
}

/**
 *  新建用户
 */
func NewUser(username string, password string, ip string, channel string, channelUID string, openID string, avatar string) *db.DBUser {
	user := &db.DBUser{
		Username: username,
		Password: password,
		Salt:     "",
		Channel:  channel,
		ChannelUID: channelUID,
		Mobile:   "",
		IP:       ip,
		OpenID:   openID,
		Status:   0,
		Avatar:   avatar,
		RegTime:  time.Now(),
		//LastLogin:time.Now(),
	}
	user.ID = db.DatabaseHandler.GenId(user)
	err := db.DatabaseHandler.Insert(user)
	if (err != nil) {
		log.Debug("add user invalid %v", err)
		exception.GameException(exception.USERNAME_EXISTS)
	}
	UserCache.SetUsernameUidMapping(user.Username, user.ID)
	UserCache.HSetUser(user.ID, user)
	return user
}

/**
 *  获取用户数据
 */
func GetUser(username string) *db.DBUser {
	uid := UserCache.GetUidByUsername(username)
	if uid == 0 {
		return nil
	}
	user := &db.DBUser{}
	UserCache.HGetUser(uid, user)
	user.ID = uid
	return user
}
