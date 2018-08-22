package internal

import (
	"aliens/module/statistics/analysis"
	"aliens/module"
	"aliens/module/statistics/conf"
)

var (
	skeleton = module.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) GetConfig() interface{} {
	return conf.Config
}

func (m *Module) GetName() string {
	return "statistics"
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	analysis.Init()
}

func (m *Module) OnDestroy() {
}