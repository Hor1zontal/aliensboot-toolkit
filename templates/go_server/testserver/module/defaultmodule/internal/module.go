package internal

import (
	"github.com/KylinHe/aliensboot/module/base"
	"e.coding.net/aliens/aliensboot_testserver/module/defaultmodule/cache"
	"e.coding.net/aliens/aliensboot_testserver/module/defaultmodule/conf"
	"e.coding.net/aliens/aliensboot_testserver/module/defaultmodule/db"
	"e.coding.net/aliens/aliensboot_testserver/module/defaultmodule/service"
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "defaultmodule"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	conf.Init()
	db.Init()
	cache.Init()
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	cache.Close()
	db.Close()
	conf.Close()
}
