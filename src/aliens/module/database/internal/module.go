package internal

import (
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

}

func (m *Module) OnDestroy() {
}