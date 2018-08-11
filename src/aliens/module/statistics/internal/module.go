package internal

import (
	"github.com/name5566/leaf/module"
	"aliens/module/statistics/analysis"
	"aliens/module/base"
)

var (
	skeleton = base.NewSkeleton()
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