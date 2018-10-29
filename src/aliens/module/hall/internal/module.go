package internal

import (
	"aliens/module/base"
	"aliens/module/hall/service"
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
	service.Init()
}

func (m *Module) OnDestroy() {
	service.Close()
}
