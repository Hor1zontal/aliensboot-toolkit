package internal

import (
	"aliens/module/cluster/cache"
	"aliens/module/cluster/dispatch"
)


type Module struct {
}

func (m *Module) IsEnable() bool {
	return true
}

func (m *Module) OnInit() {
	cache.Init()
	dispatch.Init()

}

func (m *Module) OnDestroy() {
	dispatch.Close()
	cache.Close()
}

func (s *Module) Run(closeSig chan bool) {

}
