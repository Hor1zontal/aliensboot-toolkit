package internal

import (
	"aliens/testserver/module/base"
	"aliens/testserver/module/room/service"
	"github.com/name5566/leaf/module"
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
