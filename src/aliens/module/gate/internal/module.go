package internal

import (
	"aliens/module/gate/socket"
	"aliens/module/gate/http"
	"aliens/module/gate/conf"
	"aliens/module/gate/service"
)



type Module struct {

}

func (m *Module) IsEnable() bool {
	return true
}


func (m *Module) OnInit() {
	conf.Init()
	socket.Init()
	service.Init()
	http.Init()
}

func (m *Module) OnDestroy() {
	http.Close()
	service.Close()
}

func (s *Module) Run(closeSig chan bool) {
	go socket.Skeleton.Run(closeSig)
	go socket.GateProxy.Run(closeSig)
}

