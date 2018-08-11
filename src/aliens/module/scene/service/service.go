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
    "aliens/cluster/center/service"
    "aliens/module/scene/conf"
    "aliens/cluster/center"
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
	
	if request.GetSpaceMove() != nil {
		messageRet := &protocol.SpaceMoveRet{}
		handleSpaceMove(authID, request.GetSpaceMove(), messageRet)
		response.Scene = &protocol.Response_SpaceMoveRet{messageRet}
	}
	
	if request.GetSpaceEnter() != nil {
		messageRet := &protocol.SpaceEnterRet{}
		handleSpaceEnter(authID, request.GetSpaceEnter(), messageRet)
		response.Scene = &protocol.Response_SpaceEnterRet{messageRet}
	}
	
	if request.GetSpaceLeave() != nil {
		messageRet := &protocol.SpaceLeaveRet{}
		handleSpaceLeave(authID, request.GetSpaceLeave(), messageRet)
		response.Scene = &protocol.Response_SpaceLeaveRet{messageRet}
	}
	
	if request.GetGetState() != nil {
		messageRet := &protocol.GetStateRet{}
		handleGetState(authID, request.GetGetState(), messageRet)
		response.Scene = &protocol.Response_GetStateRet{messageRet}
	}
	
	response.Code = protocol.Code_InvalidRequest
}

