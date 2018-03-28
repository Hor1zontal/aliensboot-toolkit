package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/module"
	"aliens/module/gate/conf"
	"github.com/name5566/leaf/chanrpc"
	"time"
)

//网关处理的消息管道
var skeleton = NewSkeleton()

func NewSkeleton() *module.Skeleton {
	skeleton := &module.Skeleton{
		GoLen:              conf.GoLen,
		TimerDispatcherLen: conf.TimerDispatcherLen,
		AsynCallLen:        conf.AsynCallLen,
		ChanRPCServer:      chanrpc.NewServer(conf.ChanRPCLen),
	}
	skeleton.Init()
	return skeleton
}

type Module struct {
	*gate.Gate
}

func (m *Module) IsEnable() bool {
	return conf.Config.Enable
}


func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Config.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Config.WSAddr,
		TCPAddr:         conf.Config.TCPAddr,
		HTTPTimeout:     time.Duration(conf.HTTPTimeout),
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       Processor,
		AgentChanRPC:    skeleton.ChanRPCServer,
	}
	Init()
}

func (m *Module) OnDestroy() {

}

func (s *Module) Run(closeSig chan bool) {
	go skeleton.Run(closeSig)
	s.Gate.Run(closeSig)
}

