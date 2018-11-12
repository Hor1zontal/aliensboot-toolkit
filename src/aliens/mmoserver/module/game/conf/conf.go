package conf

import (
	"aliens/aliensbot/config"
)

var Config struct {
	Service  config.ServiceConfig
	Cache    config.CacheConfig
	Database config.DBConfig
}
