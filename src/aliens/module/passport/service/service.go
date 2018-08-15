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
	"github.com/gogo/protobuf/proto"
    "aliens/log"
    "runtime/debug"
    "aliens/exception"
    "aliens/protocol/base"
    "aliens/protocol"
    "aliens/cluster/center/service"
    "aliens/module/passport/conf"
    "aliens/cluster/center"
	"aliens/chanrpc"
)

var instance service.IService = nil

func Init(chanRpc *chanrpc.Server) {
	instance = center.PublicService(conf.Config.Service, newService(chanRpc))
}

func Close() {
	center.ReleaseService(instance)
}

func newService(chanRpc *chanrpc.Server) *protocolService {
	service := &protocolService{}
	service.chanRpc = chanRpc
	service.chanRpc.Register("m", service.handle)
	return service
}

type protocolService struct {
	chanRpc *chanrpc.Server
}

func (this *protocolService) Request(request *base.Any, server base.RPCService_RequestServer) error {
	if this.chanRpc != nil {
		this.chanRpc.Call0("m", request, server)
		return nil
	}
	return nil
}


func (this *protocolService) handle(args []interface{}) {
	request := args[0].(*base.Any)
	server := args[1].(base.RPCService_RequestServer)
	requestProxy := &protocol.Request{}
	responseProxy := &protocol.Response{}
	authID := request.GetAuthId()
	defer func() {
		//处理消息异常
		if err := recover(); err != nil {
			switch err.(type) {
			case protocol.Code:
				responseProxy.Code = err.(protocol.Code)
				break
			default:
				log.Error("%v", err)
				debug.PrintStack()
				responseProxy.Code = protocol.Code_ServerException
			}
		}
		data, _ := proto.Marshal(responseProxy)
		responseProxy.Session = requestProxy.GetSession()
		server.Send(&base.Any{AuthId:authID, TypeUrl:"", Value:data})
	}()
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		exception.GameException(protocol.Code_InvalidRequest)
	}
	authID = handleRequest(requestProxy, responseProxy)
}

func handleRequest(request *protocol.Request, response *protocol.Response) int64 {
	
	if request.GetLoginRegister() != nil {
		messageRet := &protocol.LoginRegisterRet{}
		result := handleLoginRegister(request.GetLoginRegister(), messageRet)
		response.Passport = &protocol.Response_LoginRegisterRet{messageRet}
		return result
	}
	
	if request.GetLoginLogin() != nil {
		messageRet := &protocol.LoginLoginRet{}
		result := handleLoginLogin(request.GetLoginLogin(), messageRet)
		response.Passport = &protocol.Response_LoginLoginRet{messageRet}
		return result
	}
	
	if request.GetTokenLogin() != nil {
		messageRet := &protocol.TokenLoginRet{}
		result := handleTokenLogin(request.GetTokenLogin(), messageRet)
		response.Passport = &protocol.Response_TokenLoginRet{messageRet}
		return result
	}
	
	response.Code = protocol.Code_InvalidRequest
	return 0
}

