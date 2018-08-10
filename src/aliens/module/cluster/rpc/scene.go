/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/6/15
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package rpc

import (
	"aliens/module/cluster/dispatch"
	"aliens/protocol"
	"aliens/exception"
	"aliens/log"
	"aliens/protocol/base"
)

var Scene = &sceneRPCHandle{"scene"}


type sceneRPCHandle struct {
	name string
}


func (this *sceneRPCHandle) request(request *protocol.Request) *protocol.Response {
	rpcRet, err := dispatch.RPC.SyncRequest(this.name, request)
	if err != nil {
		log.Error(err)
		exception.GameException(protocol.Code_InvalidService)
	}
	any, ok := rpcRet.(*base.Any)
	if !ok {
		log.Error("invalid rpc ret data")
		exception.GameException(protocol.Code_InvalidService)
	}
	messageRet := &protocol.Response{}
	messageRet.Unmarshal(any.GetValue())
	return  messageRet
}


func (this *sceneRPCHandle) SpaceMove(request *protocol.SpaceMove) *protocol.SpaceMoveRet {
	message := &protocol.Request{
		Scene:&protocol.Request_SpaceMove{
			SpaceMove:request,
		},
	}
	messageRet := this.request(message)
	return messageRet.GetSpaceMoveRet()
}

func (this *sceneRPCHandle) SpaceEnter(request *protocol.SpaceEnter) *protocol.SpaceEnterRet {
	message := &protocol.Request{
		Scene:&protocol.Request_SpaceEnter{
			SpaceEnter:request,
		},
	}
	messageRet := this.request(message)
	return messageRet.GetSpaceEnterRet()
}

func (this *sceneRPCHandle) SpaceLeave(request *protocol.SpaceLeave) *protocol.SpaceLeaveRet {
	message := &protocol.Request{
		Scene:&protocol.Request_SpaceLeave{
			SpaceLeave:request,
		},
	}
	messageRet := this.request(message)
	return messageRet.GetSpaceLeaveRet()
}

func (this *sceneRPCHandle) GetState(request *protocol.GetState) *protocol.GetStateRet {
	message := &protocol.Request{
		Scene:&protocol.Request_GetState{
			GetState:request,
		},
	}
	messageRet := this.request(message)
	return messageRet.GetGetStateRet()
}
