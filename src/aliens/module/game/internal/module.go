package internal

import (
	"github.com/name5566/leaf/module"
	"aliens/module/base"
	"aliens/module/game/db"
	"aliens/module/game/service"
)

type Module struct {
	*module.Skeleton
}

func (m *Module) IsEnable() bool {
	return true
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	db.Init()
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	db.Close()
}