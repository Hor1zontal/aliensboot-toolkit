package internal

import (
	"aliens/module/cluster/cache"
	"aliens/module"
	"aliens/module/cluster/core"
	"aliens/module/cluster/conf"
)


type Module struct {
	*module.Skeleton
}

func (m *Module) GetName() string {
	return "cluster"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
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