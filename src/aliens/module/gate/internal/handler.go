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
	"aliens/protocol"
	"github.com/name5566/leaf/gate"
	//"aliens/module/cluster/dispatch"
	//"aliens/module/gate/conf"
	"aliens/module/gate/route"
	"aliens/module/cluster/dispatch"
	"aliens/module/cluster/constant"
)

const (
	CommandAgentNew   = "NewAgent" //新建agent
	CommandAgentClose = "CloseAgent" //关闭agent
	CommandAgentPush  = "Push"  //推送消息给agent
	CommandAgentAuth  = "Auth"  //验证agent权限
	CommandAgentMsg  = "Msg" //接受agent消息
)

func Init() {
	msg.Processor.SetByteOrder(true)
	Skeleton.RegisterChanRPC(CommandAgentNew, rpcNewAgent)
	Skeleton.RegisterChanRPC(CommandAgentClose, rpcCloseAgent)
	Skeleton.RegisterChanRPC(CommandAgentPush, agentPush)
	Skeleton.RegisterChanRPC(CommandAgentAuth, agentAuth)
	Skeleton.RegisterChanRPC(CommandAgentMsg, handleMessage)
	dispatch.MQ.RegisterConsumer(constant.SERVICE_GATE, HandlePush)
}

func Close() {
	dispatch.MQ.UNRegisterConsumer(constant.SERVICE_GATE)
}

//只处理推送消息
func HandlePush(request *protocol.Any) error {
	if request.ClientId != "" {
		request.Id = route.GetPushID(request.TypeUrl)
		Skeleton.ChanRPCServer.Go(CommandAgentPush, request.ClientId, request)
	}
	return nil
}

//授权
func agentAuth(args []interface{}) {
	network := args[0].(*network)
	networkManager.auth(network)
}

//推送消息
func agentPush(args []interface{}) {
	id := args[0].(string)
	networkManager.push(id, args[1])
}

//新的连接处理
func rpcNewAgent(args []interface{}) {
	agent := args[0].(gate.Agent)
	if agent.UserData() == nil {
		//打开缓存大小为5的收消息管道
		network := newNetwork(agent)
		agent.SetUserData(network)
		networkManager.addNetwork(network)
	}
}

//关闭连接处理
func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	networkManager.removeNetwork(a.UserData().(*network))
	a.SetUserData(nil)
}

//消息处理
func handleMessage(args []interface{}) {
	request := args[0]
	//消息的发送者
	gateAgent := args[1].(gate.Agent)
	data := gateAgent.UserData()
	switch data.(type) {
	case *network:
		data.(*network).AcceptMessage(request)
		break
	}
}

