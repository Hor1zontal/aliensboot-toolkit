package internal

import (
	"github.com/KylinHe/aliensboot/module/base"
	"e.coding.net/aliens/aliensboot_testserver/module/scene/cache"
	"e.coding.net/aliens/aliensboot_testserver/module/scene/conf"
	"e.coding.net/aliens/aliensboot_testserver/module/scene/db"
	"e.coding.net/aliens/aliensboot_testserver/module/scene/handler"
	"e.coding.net/aliens/aliensboot_testserver/module/scene/service"
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "scene"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	conf.Init()
	db.Init()
	cache.Init()
	handler.Init(m.Skeleton)
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	cache.Close()
	db.Close()
	conf.Close()
}
