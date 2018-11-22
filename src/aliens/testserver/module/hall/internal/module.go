package internal

import (
	"aliens/aliensbot/module/base"
	"aliens/testserver/module/hall/cache"
	"aliens/testserver/module/hall/conf"
	"aliens/testserver/module/hall/db"
	"aliens/testserver/module/hall/service"
	"aliens/testserver/module/hall/task"
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "hall"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	conf.Init()
	db.Init()
	cache.Init()
	task.Init(m.Skeleton)
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	cache.Close()
	db.Close()
	conf.Close()
}
