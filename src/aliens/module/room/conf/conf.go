package conf

import (
	"aliens/cluster/center/service"
	"aliens/config"
	"aliens/module/base"
	"aliens/network"
)

var configPath = base.BaseConfPath + "room/server.json"

var Config struct {
	Service    service.Config
	UDPService network.Config
}

func init() {
	config.LoadConfig(&Config, configPath) //加载服务器配置
}
