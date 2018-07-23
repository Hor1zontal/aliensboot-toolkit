package internal

import (
	"github.com/name5566/leaf/module"
	"aliens/module/base"
	"aliens/module/hall/service"
)

type Module struct {
	*module.Skeleton
}

func (m *Module) IsEnable() bool {
	return true
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	service.Init()
}

func (m *Module) OnDestroy() {
	service.Close()
}