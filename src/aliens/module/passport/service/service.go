/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/11/16
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"gok/network/service"
)

var PassportRPCService *service.GRPCService = nil

func Init() {
	service.ServiceManager.RegisterService(PassportService)

	//中心服务器订阅用户服务,能够讲查询到的用户服务地址返回给客户端
	service.SubscribeWBService(service.SERVICE_USER)

	//发布登录服务到中心服务器
	service.PublicWBService(service.SERVICE_PASSPORT)

	////配置了RPC，需要发布服务到ZK
	//if conf.Server.RPCAddress != "" && conf.Server.RPCPort != 0 {
	//	PassportRPCService = service.PublicRPCService(LocalPassportRPCService, conf.Server.RPCAddress, conf.Server.RPCPort)
	//}
}

func Close() {
	PassportService.Close()
	//LocalPassportRPCService.Close()
	if PassportRPCService != nil {
		PassportRPCService.Close()
	}
}
