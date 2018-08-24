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
	pushID := request.GetAuthId()
	if pushID > 0 {
		network.Manager.Push(pushID, request)
	}
	return nil
}
