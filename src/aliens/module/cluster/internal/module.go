package internal

import (
	"aliens/module/cluster/cache"
	"aliens/module"
	"aliens/module/cluster/core"
)


type Module struct {
	*module.Skeleton
}

func (m *Module) IsEnable() bool {
	return true
}

func (m *Module) OnInit() {
	m.Skeleton = module.NewSkeleton()
	cache.Init()
	core.Init()

}

func (m *Module) OnDestroy() {
	core.Close()
	cache.Close()
}