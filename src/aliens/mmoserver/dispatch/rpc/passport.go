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
	"aliens/aliensbot/exception"
	"aliens/aliensbot/log"
	"aliens/gok/dispatch"
	"aliens/gok/protocol"
)

var Passport = &passportRPCHandle{"passport"}

type passportRPCHandle struct {
	name string
}

func (this *passportRPCHandle) RequestNode(node string, request *protocol.Request) *protocol.Response {
	rpcRet, err := dispatch.RequestNodeMessage(this.name, node, request)
	if err != nil {
		log.Error(err)
		exception.GameException(protocol.Code_InvalidService)
	}
	return rpcRet
}

func (this *passportRPCHandle) UserRegister(node string, request *protocol.UserRegister) *protocol.UserRegisterRet {
	message := &protocol.Request{
		Passport: &protocol.Request_UserRegister{
			UserRegister: request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetUserRegisterRet()
}

func (this *passportRPCHandle) UserLogin(node string, request *protocol.UserLogin) *protocol.UserLoginRet {
	message := &protocol.Request{
		Passport: &protocol.Request_UserLogin{
			UserLogin: request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetUserLoginRet()
}

func (this *passportRPCHandle) TokenLogin(node string, request *protocol.TokenLogin) *protocol.TokenLoginRet {
	message := &protocol.Request{
		Passport: &protocol.Request_TokenLogin{
			TokenLogin: request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetTokenLoginRet()
}
