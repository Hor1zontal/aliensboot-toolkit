package internal

import (
	"aliens/aliensbot/gate"
	"aliens/aliensbot/module/base"
	"aliens/mmoserver/module/gate/conf"
	"aliens/mmoserver/module/gate/http"
	"aliens/mmoserver/module/gate/msg"
	"aliens/mmoserver/module/gate/network"
	"aliens/mmoserver/module/gate/route"
	"aliens/mmoserver/module/gate/service"
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
