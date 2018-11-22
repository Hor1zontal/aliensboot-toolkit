package internal

import (
	"aliens/aliensbot/module/base"
	"aliens/testserver/module/room/cache"
	"aliens/testserver/module/room/conf"
	"aliens/testserver/module/room/db"
	"aliens/testserver/module/room/service"
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "room"
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
