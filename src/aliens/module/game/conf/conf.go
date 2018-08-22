package conf

import (
	"aliens/cache/redis"
	"aliens/database/dbconfig"
	"aliens/cluster/center/service"
)



var Config struct {
	Service  service.Config
	Cache 	 redis.CacheConfig
	Database dbconfig.DBConfig
}
