package conf

import (
	"aliens/config"
	"aliens/cluster/center/service"
)


var configPath = "conf/aliens/hall/server.json"

var Config struct {
	Service  service.Config
	//Cache 	 redis.CacheConfig
	//Database dbconfig.DBConfig
}

func init() {
	config.LoadConfig(&Config, configPath) //加载服务器配置
}
