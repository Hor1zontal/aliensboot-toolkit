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


func (m *Module) GetName() string {
	return "database"
}

func (m *Module) GetConfig() interface{} {
	return nil
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

}

func (m *Module) OnDestroy() {
}