package internal

import (
	"aliens/module/gate/http"
	"aliens/module/gate/conf"
	"time"
	"aliens/module/gate/msg"
	"aliens/gate"
	"aliens/module/gate/service"
	"aliens/module/gate/route"
	"aliens/module/gate/network"
	"aliens/module/base"
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
	conf.Init(m.GetName())
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Config.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Config.WSAddr,
		TCPAddr:         conf.Config.TCPAddr,
		HTTPTimeout:     time.Duration(conf.HTTPTimeout),
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
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




