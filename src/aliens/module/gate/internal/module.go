package internal

import (
	"aliens/module/gate/conf"
	"aliens/module/gate/msg"
	"aliens/gate"
	"aliens/module/gate/service"
	"aliens/module/gate/route"
	"aliens/module/gate/network"
	"aliens/module/base"
	"aliens/module/gate/http"
)

var Skeleton = base.NewSkeleton()

type Module struct {
	*gate.Gate
}

func (m *Module) GetName() string {
	return "gate"
}

func (m *Module) GetConfig() interface{} {
	return &conf.Config
}

func (m *Module) OnInit() {
	//conf.Init(m.GetName())

	m.Gate = &gate.Gate{
		TcpConfig: conf.Config.TCP,
		WsConfig: conf.Config.WebSocket,
		Processor:       msg.Processor,
		AgentChanRPC:    Skeleton.ChanRPCServer,
	}
	route.Init()
	network.Init(Skeleton)
	service.Init(Skeleton.ChanRPCServer)
	http.Init()
}

func (m *Module) OnDestroy() {
	http.Close()
	service.Close()
}

func (m *Module) Run(closeSig chan bool) {
	go m.Gate.Run(closeSig)
 	Skeleton.Run(closeSig)
}




