/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/6/15
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package rpc

import (
	"aliens/gok/dispatch"
	"aliens/gok/protocol"
	"aliens/aliensbot/exception"
	"aliens/aliensbot/log"
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

func (this *gameRPCHandle) ChangeNickname(node string, request *protocol.ChangeNickname) *protocol.ChangeNicknameRet {
	message := &protocol.Request{
		Game:&protocol.Request_ChangeNickname{
			ChangeNickname:request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetChangeNicknameRet()
}
