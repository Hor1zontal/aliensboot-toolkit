package conf

import (
	"aliens/config"
)

var Config struct {
	Enable			   bool //是否启用模块
	RPCPort            int	//提供RPC服务的端口，本地启动RPC需要指定此端口启动
}


func init() {
	config.LoadConfig(&Config, "conf/aliens/service2/server.json")
}