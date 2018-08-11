package conf

import (
	"aliens/config"
	"aliens/cache/redis"
	"aliens/database/dbconfig"
	"aliens/cluster/center/service"
	"aliens/module/base"
)


var configPath = base.BaseConfPath + "game/server.json"

var Config struct {
	Service  service.Config
	Cache 	 redis.CacheConfig
	Database dbconfig.DBConfig

}

func init() {
	config.LoadConfig(&Config, configPath) //加载服务器配置
}