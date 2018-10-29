package internal

import (
	"aliens/module/base"
	"aliens/module/statistics/analysis"
	"aliens/module/statistics/conf"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
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
