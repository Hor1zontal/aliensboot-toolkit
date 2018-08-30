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
)

var Scene = &sceneRPCHandle{"scene"}


type sceneRPCHandle struct {
	name string
}


func (this *sceneRPCHandle) RequestNode(node string, request *protocol.Request) *protocol.Response {
	rpcRet, err := dispatch.RequestNodeMessage(this.name, node, request)
	if err != nil {
    	log.Error(err)
    	exception.GameException(protocol.Code_InvalidService)
    }
    return rpcRet
}


func (this *sceneRPCHandle) GetState(node string, request *protocol.GetState) *protocol.GetStateRet {
	message := &protocol.Request{
		Scene:&protocol.Request_GetState{
			GetState:request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetGetStateRet()
}

func (this *sceneRPCHandle) SpaceMove(node string, request *protocol.SpaceMove) *protocol.SpaceMoveRet {
	message := &protocol.Request{
		Scene:&protocol.Request_SpaceMove{
			SpaceMove:request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetSpaceMoveRet()
}

func (this *sceneRPCHandle) SpaceEnter(node string, request *protocol.SpaceEnter) *protocol.SpaceEnterRet {
	message := &protocol.Request{
		Scene:&protocol.Request_SpaceEnter{
			SpaceEnter:request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetSpaceEnterRet()
}

func (this *sceneRPCHandle) SpaceLeave(node string, request *protocol.SpaceLeave) *protocol.SpaceLeaveRet {
	message := &protocol.Request{
		Scene:&protocol.Request_SpaceLeave{
			SpaceLeave:request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetSpaceLeaveRet()
}
