package conf

import (
	"io/ioutil"
	"encoding/json"
	"aliens/log"
)

var Config struct {
	Enable			   bool //是否启用模块
	RPCPort            int	//提供RPC服务的端口，本地启动RPC需要指定此端口启动
}


func init() {
	data, err := ioutil.ReadFile("conf/aliens/service1/server.json")
	if err != nil {
		//log.Fatal("%v", err)
		return
	}
	err = json.Unmarshal(data, &Config)
	if err != nil {
		log.Critical("%v", err)
	}
	//if Config.RPCAddress != "" {
	//	return
	//}
	//Config.RPCAddress = util.GetAddress(Config.RPCPort)
}