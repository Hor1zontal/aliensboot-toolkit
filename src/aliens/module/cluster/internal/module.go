package internal

import (
	"aliens/cluster/center"
	"aliens/module/cluster/conf"
)


type Module struct {
}

func (m *Module) IsEnable() bool {
	return true
}

func (m *Module) OnInit() {
	center.ClusterCenter.SetLBS(conf.Config.LBS)
	center.ClusterCenter.ConnectCluster(conf.Config.ZKServers, 10, conf.Config.ZKName)
}

func (m *Module) OnDestroy() {
	center.ClusterCenter.Close()
}

func (s *Module) Run(closeSig chan bool) {

}
