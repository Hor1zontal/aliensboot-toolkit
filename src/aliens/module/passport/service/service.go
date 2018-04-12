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
	"aliens/cluster/center"
	"aliens/module/passport/conf"
)

var passportRPCService *center.GRPCService = nil

func Init() {
	passportRPCService = center.PublicGRPCService(conf.Config.Service, conf.Config.RPCPort, &passportService{})
}

func Close() {
	if passportRPCService != nil {
		passportRPCService.Close()
	}
}