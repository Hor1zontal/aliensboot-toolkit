package conf

import (
	"time"
	"aliens/config"
	"aliens/cache/redis"
	"aliens/database/dbconfig"
	"aliens/cluster/center/service"
)


var configPath = "conf/aliens/passport/server.json"

var Config struct {
	Service  service.Config
	Cache 	 redis.CacheConfig
	Database dbconfig.DBConfig

	DefaultChannelPWD string
	TokenExpireTime   int64
	HTTPAddress       string
	AppKey            string
}


func init() {
	config.LoadConfig(&Config, configPath) //加载服务器配置
	if Config.TokenExpireTime <= 0 {
		//默认过期时间一个月
		Config.TokenExpireTime = int64(30 * 24 *time.Hour)
	}
}

func GetTokenExpireTimestamp() int64 {
	return time.Now().Add(time.Duration(Config.TokenExpireTime)).Unix()
}
