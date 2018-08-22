/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/4/12
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package dispatch

import (
	"aliens/cluster/message"
	"aliens/protocol/base"
)

var serviceMapping = make(map[string]*message.RemoteService)

func Request(serviceName string, message *base.Any, hashKey string) (*base.Any, error) {
	service := allocService(serviceName)
	return service.Request(message, hashKey)
}

func RequestNode(serviceName string, serviceID string, message *base.Any) (*base.Any, error) {
	service := allocService(serviceName)
	return service.RequestNode(serviceID, message)
}

func Send(serviceName string, message *base.Any, hashKey string) error {
	service := allocService(serviceName)
	return service.Send(message, hashKey)
}

func SendNode(serviceName string, serviceID string, message *base.Any) error {
	service := allocService(serviceName)
	return service.SendNode(serviceID, message)
}

func allocService(serviceName string) *message.RemoteService {
	service := serviceMapping[serviceName]
	if service == nil {
		service = message.NewRemoteService(serviceName)
		serviceMapping[serviceName] = service
	}
	return service
}
