/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/3/24
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package message

import (
	"aliens/cluster/center"
	"aliens/protocol/base"
)

func NewRemoteService(serviceType string) *RemoteService {
	service := &RemoteService{
		serviceType: serviceType,
	}
	service.Init()
	return service
}

//远程调度服务 override IMessageService
type RemoteService struct {
	serviceType string //服务类型
}

func (this *RemoteService) Init() {
	center.ClusterCenter.SubscribeServices(this.serviceType)
}

func (this *RemoteService) HandleMessage(request *base.Any, param string) (*base.Any, error) {
	service := center.ClusterCenter.AllocService(this.serviceType, "")
	if service == nil {
		return nil, invalidServiceError
	}
	return service.Request(request)
}

func (this *RemoteService) HandleRemoteMessage(serviceID string, request *base.Any) (*base.Any, error) {
	service := center.ClusterCenter.GetService(this.serviceType, serviceID)
	if service == nil {
		return nil, invalidServiceError
	}
	return service.Request(request)
}

func (this *RemoteService) HandlePriorityRemoteMessage(serviceID string, request *base.Any) (*base.Any, error) {
	service := center.ClusterCenter.GetService(this.serviceType, serviceID)
	if service == nil {
		service = center.ClusterCenter.AllocService(this.serviceType, "")
	}
	if service == nil {
		return nil, invalidServiceError
	}
	return service.Request(request)
}

func (this *RemoteService) BroadcastAll(message *base.Any) {
	services := center.ClusterCenter.GetAllService(this.serviceType)
	if services == nil || len(services) == 0 {
		return
	}
	for _, service := range services {
		service.Request(message)
	}
	return
}

//获取消息服务类型
func (this *RemoteService) GetType() string {
	return this.serviceType
}
