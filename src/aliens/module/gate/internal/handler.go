/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/5/17
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package internal

import (
	"aliens/module/gate/msg"
	"aliens/protocol/base"
	"aliens/gate"
	"aliens/module/gate/service"
)

func init() {
	msg.Processor.SetByteOrder(true)
	Skeleton.RegisterChanRPC(gate.CommandAgentNew, newAgent)
	Skeleton.RegisterChanRPC(gate.CommandAgentClose, closeAgent)
	Skeleton.RegisterChanRPC(gate.CommandAgentMsg, handleMessage)
	//dispatch.MQ.RegisterConsumer(constant.SERVICE_GATE, HandlePush)
}


//只处理推送消息
//func HandlePush(request *base.Any) error {
//	if request.AuthId != 0 {
//		request.Id = route.GetPushID(request.TypeUrl)
//		Skeleton.ChanRPCServer.Go(CommandAgentPush, request.AuthId, request)
//	}
//	return nil
//}


//推送消息
//func agentPush(args []interface{}) {
//	networkManager.push(args[0].(int64), args[1])
//}

//新的连接处理
func newAgent(args []interface{}) {
	agent := args[0].(gate.Agent)
	if agent.UserData() == nil {
		//打开缓存大小为5的收消息管道
		network := service.NewNetwork(agent)
		agent.SetUserData(network)
		service.Manager.AddNetwork(network)
	}
}

//关闭连接处理
func closeAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	service.Manager.RemoveNetwork(a.UserData().(*service.Network))
	a.SetUserData(nil)
}

//消息处理
func handleMessage(args []interface{}) {
	request := args[0]
	//消息的发送者
	gateAgent := args[1].(gate.Agent)
	data := gateAgent.UserData()
	switch data.(type) {
	case *service.Network:
		data.(*service.Network).AcceptMessage(request.(*base.Any))
		break
	}
}

