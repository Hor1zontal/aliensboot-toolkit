package internal

import (
	"aliens/aliensbot/module/base"
	"aliens/mmoserver/module/game/conf"
	"aliens/mmoserver/module/game/db"
	"aliens/mmoserver/module/game/service"
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "game"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	db.Init()
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	db.Close()
}
