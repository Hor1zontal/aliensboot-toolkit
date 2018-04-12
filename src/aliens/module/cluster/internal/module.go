package internal

import (
	"aliens/cluster/center"
	"aliens/module/cluster/conf"
	"aliens/module/cluster/cache"
)


type Module struct {
}

func (m *Module) IsEnable() bool {
	return true
}

func (m *Module) OnInit() {
	cache.Init()
	center.ClusterCenter.ConnectCluster(conf.Config.Cluster)
}

func (m *Module) OnDestroy() {
	center.ClusterCenter.Close()
	cache.Close()
}

func (s *Module) Run(closeSig chan bool) {

}
