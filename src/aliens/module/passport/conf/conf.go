package conf

import (
	"time"
	"aliens/cache/redis"
	"aliens/database/dbconfig"
	"aliens/cluster/center/service"
)

var Config struct {
	Service  service.Config
	Cache 	 redis.CacheConfig
	Database dbconfig.DBConfig

	DefaultChannelPWD string
	TokenExpireTime   int64
	HTTPAddress       string
	AppKey            string
}

func Init() {
	if Config.TokenExpireTime <= 0 {
		//默认过期时间七天
		Config.TokenExpireTime = int64(7 * 24 *time.Hour)
	}
}

//func GetTokenExpireTimestamp() int64 {
//	return time.Now().Add(time.Duration(Config.TokenExpireTime)).Unix()
//}
