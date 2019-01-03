package internal

import (
	"github.com/KylinHe/aliensboot/module/base"
	"e.coding.net/aliens/aliensboot_testserver/module/passport/cache"
	"e.coding.net/aliens/aliensboot_testserver/module/passport/conf"
	"e.coding.net/aliens/aliensboot_testserver/module/passport/db"
	"e.coding.net/aliens/aliensboot_testserver/module/passport/service"
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "passport"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	db.Init()
	cache.Init()
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	db.Close()
	cache.Close()
}
