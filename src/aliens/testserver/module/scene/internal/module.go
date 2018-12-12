package internal

import (
	"aliens/aliensbot/module/base"
	"aliens/testserver/module/scene/cache"
	"aliens/testserver/module/scene/conf"
	"aliens/testserver/module/scene/db"
	"aliens/testserver/module/scene/handler"
	"aliens/testserver/module/scene/service"
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "scene"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	conf.Init()
	db.Init()
	cache.Init()
	handler.Init(m.Skeleton)
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	cache.Close()
	db.Close()
	conf.Close()
}
