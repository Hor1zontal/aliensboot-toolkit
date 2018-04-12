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
	center.ClusterCenter.SetLBS(conf.Config.LBS)
	center.ClusterCenter.ConnectCluster(conf.Config.ZKServers, 10, conf.Config.ZKName, conf.NodeName)
}

func (m *Module) OnDestroy() {
	center.ClusterCenter.Close()
	cache.Close()
}

func (s *Module) Run(closeSig chan bool) {

}
