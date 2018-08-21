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
    "aliens/chanrpc"
    "aliens/log"
    "runtime/debug"
    "aliens/exception"
    "aliens/protocol/base"
    "aliens/protocol"
    "aliens/cluster/center/service"
    "aliens/module/game/conf"
    "aliens/cluster/center"
)

var instance service.IService = nil

func Init(chanRpc *chanrpc.Server) {
	instance = center.PublicService(conf.Config.Service, service.NewRpcHandler(chanRpc, handle))
}

func Close() {
	center.ReleaseService(instance)
}


func handle(request *base.Any) *base.Any {
	requestProxy := &protocol.Request{}
	responseProxy := &protocol.Response{}
	response := &base.Any{}
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
		response.Value = data
	}()
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		exception.GameException(protocol.Code_InvalidRequest)
	}
	handleRequest(request.GetAuthId(), requestProxy, responseProxy)
	return response
}

func handleRequest(authID int64, request *protocol.Request, response *protocol.Response) {
	
	if request.GetGetUserInfo() != nil {
		messageRet := &protocol.GetUserInfoRet{}
		handleGetUserInfo(authID, request.GetGetUserInfo(), messageRet)
		response.Game = &protocol.Response_GetUserInfoRet{messageRet}
		return
	}
	
	if request.GetLoginRole() != nil {
		messageRet := &protocol.LoginRoleRet{}
		handleLoginRole(authID, request.GetLoginRole(), messageRet)
		response.Game = &protocol.Response_LoginRoleRet{messageRet}
		return
	}
	
	if request.GetCreateRole() != nil {
		messageRet := &protocol.CreateRoleRet{}
		handleCreateRole(authID, request.GetCreateRole(), messageRet)
		response.Game = &protocol.Response_CreateRoleRet{messageRet}
		return
	}
	
	if request.GetRemoveRole() != nil {
		messageRet := &protocol.RemoveRoleRet{}
		handleRemoveRole(authID, request.GetRemoveRole(), messageRet)
		response.Game = &protocol.Response_RemoveRoleRet{messageRet}
		return
	}
	
	response.Code = protocol.Code_InvalidRequest
}

