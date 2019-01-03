package internal

import (
	"github.com/KylinHe/aliensboot/gate"
	"github.com/KylinHe/aliensboot/module/base"
	"e.coding.net/aliens/aliensboot_testserver/module/gate/cache"
	"e.coding.net/aliens/aliensboot_testserver/module/gate/conf"
	"e.coding.net/aliens/aliensboot_testserver/module/gate/http"
	"e.coding.net/aliens/aliensboot_testserver/module/gate/msg"
	"e.coding.net/aliens/aliensboot_testserver/module/gate/network"
	"e.coding.net/aliens/aliensboot_testserver/module/gate/route"
	"e.coding.net/aliens/aliensboot_testserver/module/gate/service"
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
	cache.Init()
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
