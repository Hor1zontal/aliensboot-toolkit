package conf

import (
	"aliens/aliensbot/config"
)

var Config struct {
	Service  config.ServiceConfig //grpc
	Cache    config.CacheConfig   //redis
	Database config.DBConfig      //mongo
	Space string

}

func GetServiceName() string {
	return Config.Service.Name
}
