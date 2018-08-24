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
	"aliens/protocol"
	"aliens/exception"
	"aliens/log"
	"aliens/module/cluster/dispatch"
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

//func (this *passportRPCHandle) Request(request *protocol.Request) *protocol.Response {
//	rpcRet, err := dispatch.RPC.SyncRequest(this.name, request)
//	if err != nil {
//        log.Error(err)
//        exception.GameException(protocol.Code_InvalidService)
//    }
//    return rpcRet
//}


func (this *passportRPCHandle) LoginRegister(node string, request *protocol.LoginRegister) *protocol.LoginRegisterRet {
	message := &protocol.Request{
		Passport:&protocol.Request_LoginRegister{
			LoginRegister:request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetLoginRegisterRet()
}

func (this *passportRPCHandle) LoginLogin(node string, request *protocol.LoginLogin) *protocol.LoginLoginRet {
	message := &protocol.Request{
		Passport:&protocol.Request_LoginLogin{
			LoginLogin:request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetLoginLoginRet()
}

func (this *passportRPCHandle) TokenLogin(node string, request *protocol.TokenLogin) *protocol.TokenLoginRet {
	message := &protocol.Request{
		Passport:&protocol.Request_TokenLogin{
			TokenLogin:request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetTokenLoginRet()
}
