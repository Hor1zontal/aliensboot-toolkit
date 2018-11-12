package conf

import (
	"aliens/aliensbot/cluster/center/service"
	"aliens/aliensbot/config"
	"aliens/aliensbot/network"
	"aliens/testserver/module/base"
)

var configPath = base.BaseConfPath + "room/server.json"

var Config struct {
	Service    service.Config
	UDPService network.Config
}

func init() {
	config.LoadConfig(&Config, configPath) //加载服务器配置
}
