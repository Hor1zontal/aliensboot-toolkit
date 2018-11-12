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
	"aliens/aliensbot/chanrpc"
	"aliens/aliensbot/cluster/center"
	"aliens/aliensbot/cluster/center/service"
	"aliens/aliensbot/exception"
	"aliens/aliensbot/log"
	"aliens/aliensbot/protocol/base"
	"aliens/mmoserver/module/gate/conf"
	"aliens/mmoserver/protocol"
	"github.com/gogo/protobuf/proto"
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
	defer func() {
		exception.CatchStackDetail()
	}()
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		log.Debugf("un expect request data : %v", request)
		return nil
	}
	handleRequest(requestProxy)
	return nil
}

func handleRequest(request *protocol.Request) {

	if request.GetKickOut() != nil {
		handleKickOut(request.GetKickOut())
		return
	}

	if request.GetBindService() != nil {
		handleBindService(request.GetBindService())
		return
	}

	if request.GetPushMessage() != nil {
		handlePushMessage(request.GetPushMessage())
		return
	}

	exception.GameException(protocol.Code_InvalidRequest)
}
