/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"aliens/protocol/passport"
	"github.com/gogo/protobuf/proto"
	"aliens/protocol"
	"github.com/name5566/leaf/chanrpc"
	"aliens/log"
	"runtime/debug"
	"aliens/exception"
)

func newService(chanRpc *chanrpc.Server) *passportService {
	service := &passportService{}
	service.chanRpc = chanRpc
	service.chanRpc.Register("m", service.handle)
	return service
}

type passportService struct {
	chanRpc *chanrpc.Server
}

func (this *passportService) Request(request *protocol.Any, server protocol.RPCService_RequestServer) error {
	if this.chanRpc != nil {
		this.chanRpc.Call0("m", request, server)
		return nil
	}
	return nil
}


func (this *passportService) handle(args []interface{}) {
	request := args[0].(*protocol.Any)
	server := args[1].(protocol.RPCService_RequestServer)
	requestProxy := &passport.Request{}
	responseProxy := &passport.Response{}
	authID := request.GetAuthId()
	defer func() {
		//处理消息异常
		if err := recover(); err != nil {
			switch err.(type) {
			case passport.Code:
				responseProxy.Code = err.(passport.Code)
				break
			default:
				log.Error("%v", err)
				debug.PrintStack()
				responseProxy.Code = passport.Code_ServerException
			}
		}
		data, _ := proto.Marshal(responseProxy)
		responseProxy.Session = requestProxy.GetSession()
		server.Send(&protocol.Any{AuthId:authID, TypeUrl:"", Value:data})
	}()
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		exception.GameException(passport.Code_InvalidRequest)
	}
	authID = handleRequest(requestProxy, responseProxy)
}

func handleRequest(request *passport.Request, response *passport.Response) int64 {
	
	if request.GetLoginRegister() != nil {
		messageRet := &passport.LoginRegisterRet{}
		result := handleLoginRegister(request.GetLoginRegister(), messageRet)
		response.Response = &passport.Response_LoginRegisterRet{messageRet}
		return result
	}
	
	if request.GetLoginLogin() != nil {
		messageRet := &passport.LoginLoginRet{}
		result := handleLoginLogin(request.GetLoginLogin(), messageRet)
		response.Response = &passport.Response_LoginLoginRet{messageRet}
		return result
	}
	
	if request.GetTokenLogin() != nil {
		messageRet := &passport.TokenLoginRet{}
		result := handleTokenLogin(request.GetTokenLogin(), messageRet)
		response.Response = &passport.Response_TokenLoginRet{messageRet}
		return result
	}
	
	response.Code = passport.Code_InvalidRequest
	return 0
}

