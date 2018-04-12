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
	"time"
	"aliens/common/cache/redis"
	"aliens/module/cluster/conf"
)

var ClusterCache = &cacheManager{redisClient: &redis.RedisCacheClient{}}

type cacheManager struct {
	redisClient *redis.RedisCacheClient
}

func Init() {
	ClusterCache.Init(conf.Config.RedisAddress)
}

func Close() {
	ClusterCache.Close()
}

func (this *cacheManager) Init(redisAddress string) {
	this.Init1(redisAddress, "", 0, 0, 0)
}

func (this *cacheManager) Init1(redisAddress string, password string, maxActive int, maxIdle int, idleTimeout int) {
	if maxActive == 0 {
		maxActive = 5000
	}
	if maxIdle == 0 {
		maxIdle = 2000
	}
	if idleTimeout == 0 {
		idleTimeout = 120
	}
	this.redisClient = &redis.RedisCacheClient{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		Address:     redisAddress,
		Password:    password,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
	}
	this.redisClient.Start()
}

func (this *cacheManager) Close() {
	if this.redisClient != nil {
		this.redisClient.Close()
	}
}


