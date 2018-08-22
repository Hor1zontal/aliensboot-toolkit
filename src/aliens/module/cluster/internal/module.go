package internal

import (
	"aliens/module"
	"aliens/module/cluster/conf"
	"aliens/cluster/center"
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
	//cache.Init()
	center.ClusterCenter.ConnectCluster(conf.Config.Cluster)

}

func (m *Module) OnDestroy() {
	center.ClusterCenter.Close()
	//cache.Close()
}