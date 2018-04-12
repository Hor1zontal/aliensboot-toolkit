package internal

import (
	"aliens/module/gate/socket"
	"aliens/module/gate/http"
	"aliens/module/gate/conf"
	"aliens/module/gate/rpc"
)



type Module struct {

}

func (m *Module) IsEnable() bool {
	return true
}


func (m *Module) OnInit() {
	conf.Init()
	rpc.Init()
	socket.Init()
	http.Init()
}

func (m *Module) OnDestroy() {
	rpc.Close()
	http.Close()
}

func (s *Module) Run(closeSig chan bool) {
	go socket.Skeleton.Run(closeSig)
	go socket.GateProxy.Run(closeSig)
}

