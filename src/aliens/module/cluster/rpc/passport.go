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

var Passport = &passportRPCHandle{"passport"}


type passportRPCHandle struct {
	name string
}


func (this *passportRPCHandle) request(request *protocol.Request) *protocol.Response {
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


func (this *passportRPCHandle) LoginRegister(request *protocol.LoginRegister) *protocol.LoginRegisterRet {
	message := &protocol.Request{
		Passport:&protocol.Request_LoginRegister{
			LoginRegister:request,
		},
	}
	messageRet := this.request(message)
	return messageRet.GetLoginRegisterRet()
}

func (this *passportRPCHandle) LoginLogin(request *protocol.LoginLogin) *protocol.LoginLoginRet {
	message := &protocol.Request{
		Passport:&protocol.Request_LoginLogin{
			LoginLogin:request,
		},
	}
	messageRet := this.request(message)
	return messageRet.GetLoginLoginRet()
}

func (this *passportRPCHandle) TokenLogin(request *protocol.TokenLogin) *protocol.TokenLoginRet {
	message := &protocol.Request{
		Passport:&protocol.Request_TokenLogin{
			TokenLogin:request,
		},
	}
	messageRet := this.request(message)
	return messageRet.GetTokenLoginRet()
}
