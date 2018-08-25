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
    "aliens/exception"
    "aliens/protocol/base"
    "aliens/protocol"
    "aliens/cluster/center/service"
    "aliens/module/gate/conf"
    "aliens/cluster/center"
    "aliens/module/gate/network"
)

var instance service.IService = nil

func Init(chanRpc *chanrpc.Server) {
	instance = center.PublicService(conf.Config.Service, service.NewRpcHandler(chanRpc, handle))
}

func Close() {
	center.ReleaseService(instance)
}


func handle(request *base.Any) *base.Any {
    if request.GetGate() {
        requestProxy := &protocol.Request{}
        defer func() {
            exception.CatchStackDetail()
        }()
        error := proto.Unmarshal(request.Value, requestProxy)
        if error != nil {
            log.Debugf("un expect request data : %v", request)
            return nil
        }
        handleRequest(request.GetAuthId(), requestProxy)
    } else {
        //不是网关处理的消息即为推送消息
        network.Manager.Push(request.GetAuthId(), request)
    }
	return nil
}

func handleRequest(authID int64, request *protocol.Request) {
	
	if request.GetKickOut() != nil {
		handleKickOut(authID, request.GetKickOut())
		return
	}
	
	if request.GetBindService() != nil {
		handleBindService(authID, request.GetBindService())
		return
	}
	
	exception.GameException(protocol.Code_InvalidRequest)
}

