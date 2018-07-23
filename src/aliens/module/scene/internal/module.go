package internal

import (
	"aliens/module/scene/entity"
	"aliens/module/scene/service"
	"aliens/module/base"
	"github.com/name5566/leaf/module"
)


type Module struct {
	*module.Skeleton
}

func (m *Module) IsEnable() bool {
	return true
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	entity.Init()
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
}