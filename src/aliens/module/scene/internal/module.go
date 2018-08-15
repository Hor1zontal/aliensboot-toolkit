package internal

import (
	"aliens/module/scene/entity"
	"aliens/module/scene/service"
	"aliens/module"
)


type Module struct {
	*module.Skeleton
}

func (m *Module) IsEnable() bool {
	return true
}

func (m *Module) OnInit() {
	m.Skeleton = module.NewSkeleton()
	entity.Init()
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
}