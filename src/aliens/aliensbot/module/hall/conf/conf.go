package conf

import (
	"aliens/aliensbot/cluster/center/service"
	"aliens/aliensbot/config"
	"aliens/testserver/module/base"
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
