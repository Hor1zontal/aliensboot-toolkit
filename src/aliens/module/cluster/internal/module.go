package internal

import (
	"aliens/module/cluster/cache"
	"aliens/module/cluster/dispatch"
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
	cache.Init()
	dispatch.Init()

}

func (m *Module) OnDestroy() {
	dispatch.Close()
	cache.Close()
}