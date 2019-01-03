package internal

import (
	"github.com/KylinHe/aliensboot/module/base"
	"e.coding.net/aliens/aliensboot_testserver/module/hall/cache"
	"e.coding.net/aliens/aliensboot_testserver/module/hall/conf"
	"e.coding.net/aliens/aliensboot_testserver/module/hall/db"
	"e.coding.net/aliens/aliensboot_testserver/module/hall/service"
	"e.coding.net/aliens/aliensboot_testserver/module/hall/task"
)

type Module struct {
	*base.Skeleton
}

func (m *Module) GetName() string {
	return "hall"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	m.Skeleton = base.NewSkeleton()
	conf.Init()
	db.Init()
	cache.Init()
	task.Init(m.Skeleton)
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	cache.Close()
	db.Close()
	conf.Close()
}
