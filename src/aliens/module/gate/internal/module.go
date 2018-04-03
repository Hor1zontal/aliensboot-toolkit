package internal

import (
	"aliens/module/gate/socket"
	"aliens/module/gate/http"
	"aliens/module/gate/conf"
)



type Module struct {

}

func (m *Module) IsEnable() bool {
	return true
}


func (m *Module) OnInit() {
	conf.Init()
	socket.Init()
	http.Init()
}

func (m *Module) OnDestroy() {

}

func (s *Module) Run(closeSig chan bool) {
	go socket.Skeleton.Run(closeSig)
	socket.GateProxy.Run(closeSig)
}

