package internal

import (
	"aliens/module/cluster/cache"
	"aliens/module/cluster/dispatch"
	"github.com/name5566/leaf/module"
	"aliens/module/base"
)


type Module struct {
	*module.Skeleton
}

func (m *Module) IsEnable() bool {
	return true
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	cache.Init()
	dispatch.Init()

}

func (m *Module) OnDestroy() {
	dispatch.Close()
	cache.Close()
}