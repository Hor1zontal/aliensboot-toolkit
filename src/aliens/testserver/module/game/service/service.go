/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
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
    "aliens/aliensbot/chanrpc"
    "aliens/aliensbot/exception"
    "aliens/aliensbot/cluster/center/service"
    "aliens/aliensbot/cluster/center"
    "aliens/aliensbot/protocol/base"
    "aliens/testserver/protocol"
    "aliens/testserver/module/game/conf"

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
				exception.PrintStackDetail(err)
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
	handleRequest(request.GetAuthId(), request.GetGateId(), requestProxy, responseProxy)
	return response
}

func handleRequest(authID int64, gateID string, request *protocol.Request, response *protocol.Response) {
	
	if request.GetLoginRole() != nil {
		messageRet := &protocol.LoginRoleRet{}
		handleLoginRole(authID, gateID, request.GetLoginRole(), messageRet)
		response.Game = &protocol.Response_LoginRoleRet{messageRet}
		return
	}
	
	if request.GetChangeNickname() != nil {
		messageRet := &protocol.ChangeNicknameRet{}
		handleChangeNickname(authID, gateID, request.GetChangeNickname(), messageRet)
		response.Game = &protocol.Response_ChangeNicknameRet{messageRet}
		return
	}
	
	response.Code = protocol.Code_InvalidRequest
}

