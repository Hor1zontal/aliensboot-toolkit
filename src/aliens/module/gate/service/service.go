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
    "aliens/chanrpc"
    "aliens/protocol/base"
    "aliens/cluster/center/service"
    "aliens/module/gate/conf"
    "aliens/cluster/center"
	"aliens/log"
)

var instance service.IService = nil

func Init(chanRpc *chanrpc.Server) {
	manager.Init(chanRpc)
	instance = center.PublicService(conf.Config.Service, service.NewRpcHandler(chanRpc, handle))
}

func Close() {
	center.ReleaseService(instance)
}


func handle(request *base.Any) *base.Any {
	pushID := request.GetAuthId()
	if pushID == -1 {
		manager.broadcast(request)
	} else if pushID > 0 {
		manager.push(pushID, request)
	} else {
		log.Warnf("un expect push authID %v", pushID)
	}
	return nil
}

