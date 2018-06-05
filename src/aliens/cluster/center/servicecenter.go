/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/6/4
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package center

type ServiceCenter interface {

	ConnectCluster(config ClusterConfig)
	PublicService(service IService, unique bool) bool
	SubscribeService(serviceNames ...string)
	GetAllService(serviceName string) []IService
	IsConnect() bool
	Close()

}
