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
	"aliens/module/scene/conf"
)

var instance center.IService = nil

func Init() {
	instance = center.PublicService(conf.Config.Service, &sceneService{})
}

func Close() {
	if instance != nil {
		instance.Close()
	}
}
