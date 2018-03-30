package conf

import (
	"encoding/json"
	//"gok/log"
	"aliens/log"
	"io/ioutil"
	"time"
	"aliens/common/util"
	"aliens/config"
)

var Config struct {
	Enable            bool
	DBHost            string
	DBPort            int
	DBName            string
	DBUsername        string
	DBPassword        string
	RedisAddress      string
	RedisPassword     string
	RedisMaxActive    int
	RedisMaxIdle      int
	RedisIdleTimeout  int
	DefaultChannelPWD string
	TokenExpireTime   int64
	HTTPAddress       string
	AppKey            string

	RPCAddress        string //提供RPC服务的地址,信息需要注册到中心服务器供其他服务调用
	RPCPort           int    //提供RPC服务的端口，本地启动RPC需要指定此端口启动
}

func init() {
	config.LoadConfig(&Config, "conf/aliens/passport/server.json")
	if Config.RPCAddress == "" {
		Config.RPCAddress = util.GetAddress(Config.RPCPort)
	}
	if Config.TokenExpireTime <= 0 {
		//默认过期时间一个月
		Config.TokenExpireTime = int64(30 * 24 *time.Hour)
	}
}

func GetTokenExpireTimestamp() int64 {
	return time.Now().Add(time.Duration(Config.TokenExpireTime)).Unix()
}
