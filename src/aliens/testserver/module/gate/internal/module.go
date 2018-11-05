package internal

import (
	"aliens/aliensbot/gate"
	"aliens/aliensbot/module/base"
	"aliens/testserver/module/gate/conf"
	"aliens/testserver/module/gate/http"
	"aliens/testserver/module/gate/msg"
	"aliens/testserver/module/gate/network"
	"aliens/testserver/module/gate/route"
	"aliens/testserver/module/gate/service"
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
		TcpConfig:    conf.Config.TCP,
		WsConfig:     conf.Config.WebSocket,
		Processor:    msg.Processor,
		AgentChanRPC: Skeleton.ChanRPCServer,
	}
	route.Init()
	network.Init(Skeleton)
	service.Init(Skeleton.ChanRPCServer)
	http.Init(conf.Config.Http)
}

func (m *Module) OnDestroy() {
	http.Close()
	service.Close()
}

func (m *Module) Run(closeSig chan bool) {
	go m.Gate.Run(closeSig)
	Skeleton.Run(closeSig)
}
