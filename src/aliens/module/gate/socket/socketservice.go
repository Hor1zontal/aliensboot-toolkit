/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/4/2
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package socket

import (
	"github.com/name5566/leaf/module"
	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/gate"
	"aliens/module/gate/conf"
	"time"
	"aliens/module/gate/msg"
	"reflect"
	"aliens/protocol"
)

var GateProxy *gate.Gate = nil

//Socket网关处理的消息管道
var Skeleton = NewSkeleton()

var Processor = msg.NewMsgProcessor() //protobuf.NewProcessor()


func Init() {
	GateProxy = &gate.Gate{
		MaxConnNum:      conf.Config.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Config.WSAddr,
		TCPAddr:         conf.Config.TCPAddr,
		HTTPTimeout:     time.Duration(conf.HTTPTimeout),
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       Processor,
		AgentChanRPC:    Skeleton.ChanRPCServer,
	}

	Processor.SetByteOrder(true)
	Skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	Skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	Skeleton.RegisterChanRPC(reflect.TypeOf(&protocol.Any{}), handleMessage)
}

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

//新的连接处理
func rpcNewAgent(args []interface{}) {
	agent := args[0].(gate.Agent)
	if agent.UserData() == nil {
		//打开缓存大小为5的收消息管道
		network := newNetwork(agent)
		agent.SetUserData(network)
		networkManager.AddNetwork(network)
	}
	_ = agent
}

//关闭连接处理
func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	networkManager.RemoveNetwork(a.UserData().(*network))
	//userdata := a.UserData()
	//a.SetUserData(nil)
	_ = a
}

//消息处理
func handleMessage(args []interface{}) {
	request := args[0]
	// 消息的发送者
	gateAgent := args[1].(gate.Agent)
	userdata := gateAgent.UserData()
	switch userdata.(type) {
	case *network:
		userdata.(*network).AcceptMessage(request)
		break
	}
}
