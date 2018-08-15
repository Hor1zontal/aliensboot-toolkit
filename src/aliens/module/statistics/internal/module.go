package internal

import (
	"aliens/module/statistics/analysis"
	"aliens/module"
)

var (
	skeleton = module.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}


func (m *Module) IsEnable() bool {
	return true
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	analysis.Init()
}

func (m *Module) OnDestroy() {
}