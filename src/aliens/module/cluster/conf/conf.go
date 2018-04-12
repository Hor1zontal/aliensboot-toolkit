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
	"aliens/common/cache/redis"
	"aliens/cluster/center"
)

var configPath = "conf/aliens/cluster.json"

var Config struct {
	Cluster center.ClusterConfig
	Cache   redis.CacheConfig
}


func init() {
	config.LoadConfig(&Config, configPath)
}

