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
	requestProxy := &passport.PassportRequest{}
	responseProxy := &passport.PassportResponse{}
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
		log.Debugf("%v-%v", requestProxy, responseProxy)
		server.Send(&protocol.Any{TypeUrl:"", Value:data})
	}()
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		exception.GameException(passport.Code_InvalidRequest)
	}
	handleRequest(requestProxy, responseProxy)
}

func handleRequest(request *passport.PassportRequest, response *passport.PassportResponse) {
	
	if request.GetLoginLogin() != nil {
		messageRet := &passport.LoginLoginRet{}
		handleLoginLogin(request.GetLoginLogin(), messageRet)
		response.Response = &passport.PassportResponse_LoginLoginRet{messageRet}
		return
	}
	
	if request.GetNewInterface() != nil {
		messageRet := &passport.NewInterfaceRet{}
		handleNewInterface(request.GetNewInterface(), messageRet)
		response.Response = &passport.PassportResponse_NewInterfaceRet{messageRet}
		return
	}
	
	if request.GetLoginRegister() != nil {
		messageRet := &passport.LoginRegisterRet{}
		handleLoginRegister(request.GetLoginRegister(), messageRet)
		response.Response = &passport.PassportResponse_LoginRegisterRet{messageRet}
		return
	}
	
	response.Code = passport.Code_InvalidRequest
}

