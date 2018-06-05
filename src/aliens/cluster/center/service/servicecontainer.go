/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/6/5
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"sync"
)

func NewContainer(lbs string) *Container {
	return &Container{root:make(map[string]*serviceCategory), lbs:lbs}
}

type Container struct {
	sync.RWMutex
	root map[string]*serviceCategory //服务容器 key 服务名
	lbs string
}

//更新服务
func (this *Container) UpdateService(service IService) {
	this.Lock()
	defer this.Unlock()
	serviceName := service.GetName()
	if this.root[serviceName] == nil {
		this.root[serviceName] = NewServiceCategory(serviceName, this.lbs, "")
	}
	this.root[serviceName].updateService(service)
}

func (this *Container) UpdateServices(serviceName string, serviceIDs []IService) {
	this.Lock()
	defer this.Unlock()
	oldContainer := this.root[serviceName]
	serviceCategory := NewServiceCategory(serviceName, this.lbs, "")

	//TODO 关闭所有不可用的服务
	for _, service := range serviceIDs {
		//data, _, err := this.zkCon.Get(path + NODE_SPLIT + serviceID)
		//service := loadServiceFromData(data, serviceID, serviceName)
		//if service == nil {
		//	log.Errorf("%v unExpect service : %v", path, err)
		//	continue
		//}
		if oldContainer != nil {
			oldService := oldContainer.takeoutService(service)
			if oldService != nil {
				oldService.SetID(service.GetID())
				serviceCategory.updateService(oldService)
				continue
			}
		}
		//新服务需要连接上才能更新
		if service.Connect() {
			serviceCategory.updateService(service)
		}
	}
	this.root[serviceName] = serviceCategory
}


//根据服务类型获取一个空闲的服务节点
func (this *Container) AllocService(serviceType string) IService {
	this.RLock()
	defer this.RUnlock()
	//TODO 后续要优化，考虑负载、空闲等因素
	serviceCategory := this.root[serviceType]
	if serviceCategory == nil {
		return nil
	}
	return serviceCategory.allocService()
}

//
func (this *Container) GetMasterService(serviceType string) IService {
	this.RLock()
	defer this.RUnlock()
	serviceCategory := this.root[serviceType]
	if serviceCategory == nil {
		return nil
	}
	return serviceCategory.getMaster()
}

func (this *Container) GetService(serviceType string, serviceID string) IService {
	this.RLock()
	defer this.RUnlock()
	serviceCategory := this.root[serviceType]
	if serviceCategory == nil {
		return nil
	}
	return serviceCategory.services[serviceID]
	////节点没有取第一个
	//if (service == nil) {
	//	serviceCategory.allocService()
	//}
	//return service
}

func (this *Container) GetAllService(serviceType string) []IService {
	this.RLock()
	defer this.RUnlock()
	serviceCategory := this.root[serviceType]
	if serviceCategory == nil {
		return nil
	}
	return serviceCategory.getAllService()
	////节点没有取第一个
	//if (service == nil) {
	//	serviceCategory.allocService()
	//}
	//return service
}

func (this *Container) GetServiceInfo(serviceType string) []string {
	this.RLock()
	defer this.RUnlock()
	serviceCategory := this.root[serviceType]
	if serviceCategory == nil {
		return nil
	}
	return serviceCategory.getNodes()
}




