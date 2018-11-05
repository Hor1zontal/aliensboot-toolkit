package internal

import (
	"aliens/testserver/module"
	"aliens/testserver/module/scene/conf"
	"aliens/testserver/module/scene/core"
	"aliens/testserver/module/scene/service"
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
