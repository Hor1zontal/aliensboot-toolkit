package internal

import (
	"aliens/module/passport/cache"
	"aliens/module/passport/db"
	"aliens/module/passport/service"
	"github.com/name5566/leaf/module"
	"aliens/module/base"
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
	cache.Init()
	service.Init(m.ChanRPCServer)
}

func (m *Module) OnDestroy() {
	service.Close()
	db.Close()
	cache.Close()

}