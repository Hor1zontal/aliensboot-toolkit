package internal

import (
	"aliens/module/passport/cache"
	"aliens/module/passport/db"
	"aliens/module/passport/service"
	"aliens/module/passport/conf"
	"aliens/module/base"
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "passport"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	db.Init()
	cache.Init()
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	db.Close()
	cache.Close()
}