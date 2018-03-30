package internal

import (
	"gok/passportserver/cache"
	"gok/passportserver/conf"
	"gok/passportserver/db"
	"gok/passportserver/service"
	"gok/passportserver/service/http"
)

type Module struct {
}

func (m *Module) IsEnable() bool {
	return conf.Server.Enable
}

func (m *Module) OnInit() {
	db.Init()
	cache.Init()
	service.Init()
	http.Init()
}

func (m *Module) OnDestroy() {
	http.Close()
	service.Close()
	db.Close()
	cache.Close()

}

func (s *Module) Run(closeSig chan bool) {

}
