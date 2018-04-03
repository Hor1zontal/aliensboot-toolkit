package internal

import (
	"aliens/module/passport/cache"
	"aliens/module/passport/db"
	"aliens/module/passport/service"
)

type Module struct {
}

func (m *Module) IsEnable() bool {
	return true
}

func (m *Module) OnInit() {
	db.Init()
	cache.Init()
	service.Init()
}

func (m *Module) OnDestroy() {
	service.Close()
	db.Close()
	cache.Close()

}

func (s *Module) Run(closeSig chan bool) {

}
