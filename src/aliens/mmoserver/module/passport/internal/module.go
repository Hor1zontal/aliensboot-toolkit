package internal

import (
	"aliens/aliensbot/module/base"
	"aliens/mmoserver/module/passport/cache"
	"aliens/mmoserver/module/passport/conf"
	"aliens/mmoserver/module/passport/db"
	"aliens/mmoserver/module/passport/service"
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
