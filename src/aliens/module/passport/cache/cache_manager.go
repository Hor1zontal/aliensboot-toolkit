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
	"aliens/cache/redis"
	"aliens/module/passport/conf"
	"aliens/module/passport/db"
	"time"
	"aliens/log"
	"aliens/protocol/passport"
	"aliens/exception"
)

var PassportCache = &cacheManager{redisClient: &redis.RedisCacheClient{}}

type cacheManager struct {
	redisClient *redis.RedisCacheClient
}

func Init() {
	PassportCache.Init(conf.Config.Cache)

	//其他服务器加载了缓存就不需要再加载到缓存了
	if PassportCache.SetNX(FLAG_LOADUSER, 1) {
		var users []*passport.User
		db.DatabaseHandler.QueryAll(&passport.User{}, &users)
		for _, user := range users {
			PassportCache.SetUsernameUidMapping(user.GetUsername(), user.GetId())
			PassportCache.HSetUser(user.GetId(), user)
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
func NewUser(username string, password string, ip string, channel string, channelUID string, openID string, avatar string) *passport.User {
	user := &passport.User{
		Username: username,
		Password: password,
		Salt:     "",
		Channel:  channel,
		Channeluid: channelUID,
		Mobile:   "",
		Ip:       ip,
		Openid:   openID,
		Status:   0,
		Avatar:   avatar,
		RegTime:  time.Now(),
		//LastLogin:time.Now(),
	}
	uid, err := db.DatabaseHandler.GenTimestampId(user)
	if err != nil {
		log.Debugf("insert user error %v", err)
		exception.GameException(exception.DATABASE_ERROR)
	}
	user.Id = uid
	err1 := db.DatabaseHandler.Insert(user)
	if err1 != nil {
		log.Debugf("insert user error %v", err1)
		exception.GameException(exception.DATABASE_ERROR)
	}
	PassportCache.SetUsernameUidMapping(user.Username, user.GetId())
	PassportCache.HSetUser(user.GetId(), user)
	return user
}

/**
 *  获取用户数据
 */
func GetUser(username string) *passport.User {
	uid := PassportCache.GetUidByUsername(username)
	if uid == 0 {
		return nil
	}
	user := &passport.User{}
	PassportCache.HGetUser(uid, user)
	user.Id = uid
	return user
}
