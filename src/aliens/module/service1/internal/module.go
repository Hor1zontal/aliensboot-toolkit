package internal

import (
	"aliens/module/service1/service"
	"aliens/module/service1/conf"
)

type Module struct {
}

func (m *Module) IsEnable() bool {
	return conf.Config.Enable
}


func (m *Module) OnInit() {
	service.Init()
}

func (m *Module) OnDestroy() {
	service.Close()
}

func (s *Module) Run(closeSig chan bool) {

}

