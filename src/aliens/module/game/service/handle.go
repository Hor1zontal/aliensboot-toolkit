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
	"aliens/protocol/game"
	"github.com/gogo/protobuf/proto"
	"aliens/protocol"
	"github.com/name5566/leaf/chanrpc"
	"aliens/log"
	"runtime/debug"
	"aliens/exception"
)

func newService(chanRpc *chanrpc.Server) *gameService {
	service := &gameService{}
	service.chanRpc = chanRpc
	service.chanRpc.Register("m", service.handle)
	return service
}

type gameService struct {
	chanRpc *chanrpc.Server
}

func (this *gameService) Request(request *protocol.Any, server protocol.RPCService_RequestServer) error {
	if this.chanRpc != nil {
		this.chanRpc.Call0("m", request, server)
		return nil
	}
	return nil
}


func (this *gameService) handle(args []interface{}) {
	request := args[0].(*protocol.Any)
	server := args[1].(protocol.RPCService_RequestServer)
	requestProxy := &game.Request{}
	responseProxy := &game.Response{}
	defer func() {
		//处理消息异常
		if err := recover(); err != nil {
			switch err.(type) {
			case game.Code:
				responseProxy.Code = err.(game.Code)
				break
			default:
				log.Error("%v", err)
				debug.PrintStack()
				responseProxy.Code = game.Code_ServerException
			}
		}
		data, _ := proto.Marshal(responseProxy)
		responseProxy.Session = requestProxy.GetSession()
		server.Send(&protocol.Any{Value:data})
	}()
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		exception.GameException(game.Code_InvalidRequest)
	}
	handleRequest(request.GetAuthId(), requestProxy, responseProxy)
}

func handleRequest(authID int64, request *game.Request, response *game.Response) {
	
	if request.GetGetUserInfo() != nil {
		messageRet := &game.GetUserInfoRet{}
		handleGetUserInfo(authID, request.GetGetUserInfo(), messageRet)
		response.Response = &game.Response_GetUserInfoRet{messageRet}
		return
	}
	
	if request.GetLoginRole() != nil {
		messageRet := &game.LoginRoleRet{}
		handleLoginRole(authID, request.GetLoginRole(), messageRet)
		response.Response = &game.Response_LoginRoleRet{messageRet}
		return
	}
	
	if request.GetCreateRole() != nil {
		messageRet := &game.CreateRoleRet{}
		handleCreateRole(authID, request.GetCreateRole(), messageRet)
		response.Response = &game.Response_CreateRoleRet{messageRet}
		return
	}
	
	if request.GetRemoveRole() != nil {
		messageRet := &game.RemoveRoleRet{}
		handleRemoveRole(authID, request.GetRemoveRole(), messageRet)
		response.Response = &game.Response_RemoveRoleRet{messageRet}
		return
	}
	
	response.Code = game.Code_InvalidRequest
}

