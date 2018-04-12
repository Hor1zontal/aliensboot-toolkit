/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/3/29
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package cache

import (
	"aliens/common/cache/redis"
	"aliens/module/passport/conf"
	"aliens/module/passport/db"
	"time"
	"aliens/log"
)

var PassportCache = &cacheManager{redisClient: &redis.RedisCacheClient{}}

type cacheManager struct {
	redisClient *redis.RedisCacheClient
}

func Init() {
	PassportCache.Init(conf.Config.Cache)

	//其他服务器加载了缓存就不需要再加载到缓存了
	if PassportCache.SetNX(FLAG_LOADUSER, 1) {
		var users []*db.DBUser
		db.DatabaseHandler.QueryAll(&db.DBUser{}, &users)
		for _, user := range users {
			PassportCache.SetUsernameUidMapping(user.Username, user.ID)
			PassportCache.HSetUser(user.ID, user)
		}
	}
}

func Close() {
	PassportCache.Close()
}

func (this *cacheManager) Init(config redis.CacheConfig) {
	this.redisClient = redis.NewRedisClient(config)
	this.redisClient.Start()
}

func (this *cacheManager) Close() {
	if this.redisClient != nil {
		this.redisClient.Close()
	}
}

func (this *cacheManager) SetNX(key string, value interface{}) bool {
	return this.redisClient.SetNX(key, value)
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
	user.ID = db.DatabaseHandler.GenTimestampId(user)
	err := db.DatabaseHandler.Insert(user)
	if err != nil {
		log.Debug("add user invalid %v", err)
		//exception.GameException(exception.USERNAME_EXISTS)
	}
	PassportCache.SetUsernameUidMapping(user.Username, user.ID)
	PassportCache.HSetUser(user.ID, user)
	return user
}

/**
 *  获取用户数据
 */
func GetUser(username string) *db.DBUser {
	uid := PassportCache.GetUidByUsername(username)
	if uid == 0 {
		return nil
	}
	user := &db.DBUser{}
	PassportCache.HGetUser(uid, user)
	user.ID = uid
	return user
}
