package internal

import (
	"github.com/name5566/leaf/gate"
	"aliens/cluster/message"
	"aliens/module/gate/msg"
	"reflect"
	"github.com/gogo/protobuf/types"
)

var router = make(map[uint16]message.IMessageService)

var Processor = msg.NewMsgProcessor() //protobuf.NewProcessor()

func Init() {
	Processor.SetByteOrder(true)
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC(reflect.TypeOf(&types.Any{}), handleMessage)
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