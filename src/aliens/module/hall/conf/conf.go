package conf

import (
	"aliens/cluster/center/service"
	"aliens/config"
	"aliens/module/base"
)

var configPath = base.BaseConfPath + "hall/server.json"

var Config struct {
	Service service.Config
	//Cache 	 redis.CacheConfig
	//Database dbconfig.DBConfig
}

func init() {
	config.LoadConfig(&Config, configPath) //加载服务器配置
}
