package internal

import (
	"aliens/module/gate/http"
	"aliens/module/gate/conf"
	"time"
	"github.com/name5566/leaf/gate"
	"aliens/module/gate/msg"
	"aliens/module/base"
)

var Skeleton = base.NewSkeleton()

type Module struct {
	*gate.Gate
}

func (m *Module) IsEnable() bool {
	return true
}


func (m *Module) OnInit() {
	conf.Init()
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
	Init()
	http.Init()
}

func (m *Module) OnDestroy() {
	http.Close()
	Close()
}

func (m *Module) Run(closeSig chan bool) {
	go m.Gate.Run(closeSig)
	Skeleton.Run(closeSig)
}




