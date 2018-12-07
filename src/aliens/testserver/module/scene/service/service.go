// Code generated by aliensbot. DO NOT EDIT.
// source: module scene
package service

import (
	"aliens/testserver/constant"
	"github.com/gogo/protobuf/proto"
    "aliens/aliensbot/chanrpc"
    "aliens/aliensbot/exception"
    "aliens/aliensbot/cluster/center/service"
    "aliens/aliensbot/cluster/center"
    "aliens/aliensbot/protocol/base"
    "aliens/testserver/protocol"
    "aliens/testserver/module/scene/conf"

)

var instance service.IService = nil

func Init(chanRpc *chanrpc.Server) {
	instance = center.PublicService(conf.Config.Service, service.NewRpcHandler(chanRpc, handle))
}

func Close() {
	center.ReleaseService(instance)
}


func handle(request *base.Any) (response *base.Any) {
	if request.Id == constant.MsgOffline {
		handleOffline(request.GetAuthId())
	}
	requestProxy := &protocol.Request{}
	responseProxy := &protocol.Response{}
	response = &base.Any{}
	isResponse := false
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
		if !isResponse {
            return
        }
		data, _ := proto.Marshal(responseProxy)
		responseProxy.Session = requestProxy.GetSession()
		response.Value = data
	}()
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		exception.GameException(protocol.Code_InvalidRequest)
	}
	isResponse = handleRequest(request.GetAuthId(), request.GetGateId(), requestProxy, responseProxy)
	return
}

func handleRequest(authID int64, gateID string, request *protocol.Request, response *protocol.Response) bool {
    if request.GetEntityCall() != nil {
    	handleEntityCall(authID, gateID, request.GetEntityCall())
    	return false
    }
    
    if request.GetLoginScene() != nil {
    	handleLoginScene(authID, gateID, request.GetLoginScene())
    	return false
    }
    
	response.Code = protocol.Code_InvalidRequest
	return true
}

