package conf

import (
	"aliens/config"
)

var Config struct {
	Service  config.ServiceConfig
	Cache    config.CacheConfig
	Database config.DBConfig
}
