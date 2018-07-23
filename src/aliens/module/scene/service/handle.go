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
	"aliens/protocol/scene"
	"github.com/gogo/protobuf/proto"
	"aliens/protocol"
	"github.com/name5566/leaf/chanrpc"
	"aliens/log"
	"runtime/debug"
	"aliens/exception"
)

func newService(chanRpc *chanrpc.Server) *sceneService {
	service := &sceneService{}
	service.chanRpc = chanRpc
	service.chanRpc.Register("m", service.handle)
	return service
}

type sceneService struct {
	chanRpc *chanrpc.Server
}

func (this *sceneService) Request(request *protocol.Any, server protocol.RPCService_RequestServer) error {
	if this.chanRpc != nil {
		this.chanRpc.Call0("m", request, server)
		return nil
	}
	return nil
}


func (this *sceneService) handle(args []interface{}) {
	request := args[0].(*protocol.Any)
	server := args[1].(protocol.RPCService_RequestServer)
	requestProxy := &scene.SceneRequest{}
	responseProxy := &scene.SceneResponse{}
	defer func() {
		//处理消息异常
		if err := recover(); err != nil {
			switch err.(type) {
			case scene.Code:
				responseProxy.Code = err.(scene.Code)
				break
			default:
				log.Error("%v", err)
				debug.PrintStack()
				responseProxy.Code = scene.Code_ServerException
			}
		}
		data, _ := proto.Marshal(responseProxy)
		responseProxy.Session = requestProxy.GetSession()
		log.Debugf("%v-%v", requestProxy, responseProxy)
		server.Send(&protocol.Any{TypeUrl:"", Value:data})
	}()
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		exception.GameException(scene.Code_InvalidRequest)
	}
	handleRequest(requestProxy, responseProxy)
}

func handleRequest(request *scene.SceneRequest, response *scene.SceneResponse) {
	
	if request.GetSpaceMove() != nil {
		messageRet := &scene.SpaceMoveRet{}
		handleSpaceMove(request.GetSpaceMove(), messageRet)
		response.Response = &scene.SceneResponse_SpaceMoveRet{messageRet}
		return
	}
	
	if request.GetSpaceEnter() != nil {
		messageRet := &scene.SpaceEnterRet{}
		handleSpaceEnter(request.GetSpaceEnter(), messageRet)
		response.Response = &scene.SceneResponse_SpaceEnterRet{messageRet}
		return
	}
	
	if request.GetSpaceLeave() != nil {
		messageRet := &scene.SpaceLeaveRet{}
		handleSpaceLeave(request.GetSpaceLeave(), messageRet)
		response.Response = &scene.SceneResponse_SpaceLeaveRet{messageRet}
		return
	}
	
	if request.GetGetState() != nil {
		messageRet := &scene.GetStateRet{}
		handleGetState(request.GetGetState(), messageRet)
		response.Response = &scene.SceneResponse_GetStateRet{messageRet}
		return
	}
	
	response.Code = scene.Code_InvalidRequest
}

