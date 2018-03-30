package internal

import (
	"aliens/module/scene/entity"
	"aliens/module/scene/service"
)


type Module struct {
}

func (m *Module) IsEnable() bool {
	return true
}

func (m *Module) OnInit() {
	entity.Init()
	service.Init()
}

func (m *Module) OnDestroy() {

}

func (s *Module) Run(closeSig chan bool) {

}
