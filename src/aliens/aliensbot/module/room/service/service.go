/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/11/16
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"aliens/aliensbot/cluster/center"
	"aliens/aliensbot/cluster/center/service"
	"aliens/testserver/module/room/conf"
)

var instance service.IService = nil

func Init() {
	instance = center.PublicService(conf.Config.Service, &roomService{})
}

func Close() {
	center.ReleaseService(instance)
}