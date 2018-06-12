package conf

import (
	"aliens/config"
	"aliens/cluster/center/service"
	"aliens/network"
)


var configPath = "conf/aliens/room/server.json"

var Config struct {
	Service  service.Config
	UDPService network.Config
}

func init() {
	config.LoadConfig(&Config, configPath) //加载服务器配置
}