package internal

import (
	"aliens/module/scene/service"
	"aliens/module"
	"aliens/module/scene/conf"
	"aliens/module/scene/core"
)


type Module struct {
	*module.Skeleton
}

func (m *Module) GetName() string {
	return "scene"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}


func (m *Module) OnInit() {
	m.Skeleton = module.NewSkeleton()
	core.Init()
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
}