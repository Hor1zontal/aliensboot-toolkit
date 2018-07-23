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
	"aliens/cluster/center/service"
	"aliens/module/hall/conf"
	"aliens/cluster/center"
)

var instance service.IService = nil

func Init() {
	instance = center.PublicService(conf.Config.Service, &hallService{})
}

func Close() {
	center.ReleaseService(instance)
}
