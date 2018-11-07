package internal

import (
	"aliens/aliensbot/module/base"
	"aliens/testserver/module/defaultmodule/conf"
	"aliens/testserver/module/defaultmodule/db"
	"aliens/testserver/module/defaultmodule/service"
	"aliens/testserver/module/defaultmodule/cache"
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "defaultmodule"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	conf.Init()
	db.Init()
	cache.Init()
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	cache.Close()
	db.Close()
	conf.Close()
}
