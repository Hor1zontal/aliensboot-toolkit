package conf

import (
	"github.com/KylinHe/aliensboot/config"
)

var Config struct {
	Service  config.ServiceConfig
	Cache    config.CacheConfig
	Database config.DBConfig
}
