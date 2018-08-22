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
	"aliens/cache/redis"
	"aliens/cluster/center"
	"aliens/mq"
)

var Config struct {
	Cluster center.ClusterConfig
	Cache   redis.CacheConfig
	MQ mq.Config
}
