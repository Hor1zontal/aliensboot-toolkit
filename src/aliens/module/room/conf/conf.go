package conf

import (
	"aliens/config"
	"aliens/cluster/center/service"
	"aliens/network"
	"aliens/module/base"
)


var configPath = base.BaseConfPath + "room/server.json"

var Config struct {
	Service  service.Config
	UDPService network.Config
}

func init() {
	config.LoadConfig(&Config, configPath) //加载服务器配置
}