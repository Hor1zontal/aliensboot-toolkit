package internal

import (
	"aliens/module/cluster/conf"
	"aliens/cluster/center"
	"aliens/module/cluster/cache"
	"aliens/module/base"
)

var Skeleton = base.NewSkeleton()

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "cluster"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}


func (m *Module) OnInit() {
	m.Skeleton = Skeleton
	cache.Init()
	center.ClusterCenter.ConnectCluster(conf.Config.Cluster)

}

func (m *Module) OnDestroy() {
	center.ClusterCenter.Close()
	cache.Close()
}