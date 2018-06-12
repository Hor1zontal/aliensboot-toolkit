package internal

import (
	"aliens/module/base"
	"github.com/name5566/leaf/module"
	"aliens/module/room/service"
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
	service.Init()
}

func (m *Module) OnDestroy() {
	service.Close()
}