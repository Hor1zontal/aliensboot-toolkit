package internal

import (
	"aliens/module/passport/cache"
	"aliens/module/passport/db"
	"aliens/module/passport/service"
	"aliens/module"
	"aliens/module/passport/conf"
)

type Module struct {
	*module.Skeleton
}

func (m *Module) GetName() string {
	return "passport"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = module.NewSkeleton()
	db.Init()
	cache.Init()
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	db.Close()
	cache.Close()
}