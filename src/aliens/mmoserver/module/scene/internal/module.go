package internal

import (
	"aliens/aliensbot/module/base"
	"aliens/mmoserver/module/scene/conf"
	"aliens/mmoserver/module/scene/core"
	"aliens/mmoserver/module/scene/service"
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
	core.Init()
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
}
