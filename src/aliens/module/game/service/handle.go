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
    "github.com/name5566/leaf/chanrpc"
    "aliens/log"
    "runtime/debug"
    "aliens/exception"
    "aliens/protocol/base"
    "aliens/protocol"
)

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
		server.Send(&base.Any{Value:data})
	}()
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		exception.GameException(protocol.Code_InvalidRequest)
	}
	handleRequest(request.GetAuthId(), requestProxy, responseProxy)
}

func handleRequest(authID int64, request *protocol.Request, response *protocol.Response) {
	
	if request.GetLoginRole() != nil {
		messageRet := &protocol.LoginRoleRet{}
		handleLoginRole(authID, request.GetLoginRole(), messageRet)
		response.Game = &protocol.Response_LoginRoleRet{messageRet}
	}
	
	if request.GetCreateRole() != nil {
		messageRet := &protocol.CreateRoleRet{}
		handleCreateRole(authID, request.GetCreateRole(), messageRet)
		response.Game = &protocol.Response_CreateRoleRet{messageRet}
	}
	
	if request.GetRemoveRole() != nil {
		messageRet := &protocol.RemoveRoleRet{}
		handleRemoveRole(authID, request.GetRemoveRole(), messageRet)
		response.Game = &protocol.Response_RemoveRoleRet{messageRet}
	}
	
	if request.GetGetUserInfo() != nil {
		messageRet := &protocol.GetUserInfoRet{}
		handleGetUserInfo(authID, request.GetGetUserInfo(), messageRet)
		response.Game = &protocol.Response_GetUserInfoRet{messageRet}
	}
	
	response.Code = protocol.Code_InvalidRequest
}

