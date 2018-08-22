package internal

import (
	"aliens/module/game/db"
	"aliens/module/game/service"
	"aliens/module"
	"aliens/module/game/conf"
)

type Module struct {
	*module.Skeleton
}

func (m *Module) GetName() string {
	return "game"
}

func (m *Module) GetConfig() interface{} {
	return conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = module.NewSkeleton()
	db.Init()
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	db.Close()
}