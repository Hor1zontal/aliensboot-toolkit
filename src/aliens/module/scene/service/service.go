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
    "aliens/module/scene/conf"
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
	
	if request.GetSpaceLeave() != nil {
		messageRet := &protocol.SpaceLeaveRet{}
		handleSpaceLeave(authID, request.GetSpaceLeave(), messageRet)
		response.Scene = &protocol.Response_SpaceLeaveRet{messageRet}
		return
	}
	
	if request.GetGetState() != nil {
		messageRet := &protocol.GetStateRet{}
		handleGetState(authID, request.GetGetState(), messageRet)
		response.Scene = &protocol.Response_GetStateRet{messageRet}
		return
	}
	
	if request.GetSpaceMove() != nil {
		messageRet := &protocol.SpaceMoveRet{}
		handleSpaceMove(authID, request.GetSpaceMove(), messageRet)
		response.Scene = &protocol.Response_SpaceMoveRet{messageRet}
		return
	}
	
	if request.GetSpaceEnter() != nil {
		messageRet := &protocol.SpaceEnterRet{}
		handleSpaceEnter(authID, request.GetSpaceEnter(), messageRet)
		response.Scene = &protocol.Response_SpaceEnterRet{messageRet}
		return
	}
	
	response.Code = protocol.Code_InvalidRequest
}

