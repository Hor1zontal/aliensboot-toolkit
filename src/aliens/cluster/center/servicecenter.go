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
	GetNodeID() string
	ConnectCluster(config ClusterConfig)
	PublicService(service IService, unique bool) bool
	SubscribeServices(serviceNames ...string)
	GetAllService(serviceName string) []IService
	GetService(serviceName string, serviceID string) IService
	AllocService(serviceName string) IService

	IsConnect() bool
	Close()
}
