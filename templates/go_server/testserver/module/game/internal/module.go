package internal

import (
	"github.com/KylinHe/aliensboot/module/base"
	"e.coding.net/aliens/aliensboot_testserver/module/game/conf"
	"e.coding.net/aliens/aliensboot_testserver/module/game/db"
	"e.coding.net/aliens/aliensboot_testserver/module/game/service"
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "game"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
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
