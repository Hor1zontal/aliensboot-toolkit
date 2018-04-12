/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/8/4
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package conf

import (
	"aliens/config"
	"gopkg.in/mgo.v2/bson"
)

var Config struct {
	ID 		  string   //节点id
	ZKServers []string //集群中心服务器地址
	ZKName    string   //集群名称，不用业务使用不同的集群
	LBS       string   //负载均衡策略  polling 轮询
	RedisAddress string //集群缓存服务器地址
}

var NodeName = bson.NewObjectId().Hex()


func init() {
	config.LoadConfig(&Config, "conf/aliens/cluster.json")
	if Config.ID != "" {
		NodeName = Config.ID
	}
}

