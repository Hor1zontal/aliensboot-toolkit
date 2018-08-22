package internal

import (
	"aliens/module/scene/entity"
	"aliens/module/scene/service"
	"aliens/module"
	"aliens/module/scene/conf"
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
	entity.Init()
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
}