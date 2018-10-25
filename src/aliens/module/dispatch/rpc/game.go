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

var Game = &gameRPCHandle{"game"}


type gameRPCHandle struct {
	name string
}


func (this *gameRPCHandle) RequestNode(node string, request *protocol.Request) *protocol.Response {
	rpcRet, err := dispatch.RequestNodeMessage(this.name, node, request)
	if err != nil {
    	log.Error(err)
    	exception.GameException(protocol.Code_InvalidService)
    }
    return rpcRet
}


func (this *gameRPCHandle) LoginRole(node string, request *protocol.LoginRole) *protocol.LoginRoleRet {
	message := &protocol.Request{
		Game:&protocol.Request_LoginRole{
			LoginRole:request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetLoginRoleRet()
}

func (this *gameRPCHandle) CreateRole(node string, request *protocol.CreateRole) *protocol.CreateRoleRet {
	message := &protocol.Request{
		Game:&protocol.Request_CreateRole{
			CreateRole:request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetCreateRoleRet()
}

func (this *gameRPCHandle) RemoveRole(node string, request *protocol.RemoveRole) *protocol.RemoveRoleRet {
	message := &protocol.Request{
		Game:&protocol.Request_RemoveRole{
			RemoveRole:request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetRemoveRoleRet()
}

func (this *gameRPCHandle) GetUserInfo(node string, request *protocol.GetUserInfo) *protocol.GetUserInfoRet {
	message := &protocol.Request{
		Game:&protocol.Request_GetUserInfo{
			GetUserInfo:request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetGetUserInfoRet()
}
