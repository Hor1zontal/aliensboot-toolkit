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


func (this *passportRPCHandle) C2S_UserRegister(node string, request *protocol.C2S_UserRegister) *protocol.S2C_UserRegister {
	message := &protocol.Request{
		Passport:&protocol.Request_C2S_UserRegister{
			C2S_UserRegister:request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetS2C_UserRegister()
}

func (this *passportRPCHandle) C2S_UserLogin(node string, request *protocol.C2S_UserLogin) *protocol.S2C_UserLogin {
	message := &protocol.Request{
		Passport:&protocol.Request_C2S_UserLogin{
			C2S_UserLogin:request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetS2C_UserLogin()
}

func (this *passportRPCHandle) C2S_TokenLogin(node string, request *protocol.C2S_TokenLogin) *protocol.S2C_TokenLogin {
	message := &protocol.Request{
		Passport:&protocol.Request_C2S_TokenLogin{
			C2S_TokenLogin:request,
		},
	}
	messageRet := this.RequestNode(node, message)
	return messageRet.GetS2C_TokenLogin()
}
