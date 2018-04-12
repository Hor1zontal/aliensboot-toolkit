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
	"aliens/common/cache"
	"time"
)


type cacheManager struct {
	*cache.RedisCacheClient
}

func (this *cacheManager) Init(redisAddress string) {
	this.Init1(redisAddress, "", 0, 0, 0)
}

func (this *cacheManager) Init1(redisAddress string, password string, maxActive int, maxIdle int, idleTimeout int) {
	if (maxActive == 0) {
		maxActive = 5000
	}
	if (maxIdle == 0) {
		maxIdle = 2000
	}
	if (idleTimeout == 0) {
		idleTimeout = 120
	}
	redisClient := &cache.RedisCacheClient{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		Address:     redisAddress,
		Password:    password,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
	}
	redisClient.Start()
	this.RedisCacheClient = redisClient
}

func (this *cacheManager) Close() {
	if this.RedisCacheClient != nil {
		this.RedisCacheClient.Close()
	}
}

func (this *cacheManager) SetNX(key string, value interface{}) bool {
	return this.RedisCacheClient.SetNX(key, value)
}

