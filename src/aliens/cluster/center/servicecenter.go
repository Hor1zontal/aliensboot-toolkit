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

import (
	"aliens/cluster/center/service"
	"aliens/log"
)

var ClusterCenter ServiceCenter = &ZKServiceCenter{} //服务调度中心


func PublicService(config service.ServiceConfig, handler interface{}) service.IService {
	if !ClusterCenter.IsConnect() {
		log.Fatal(config.Name + " cluster center is not connected")
		return nil
	}
	service := service.NewService(config)
	service.SetID(ClusterCenter.GetNodeID())
	if service == nil {
		log.Fatalf( "un expect service protocol %v", config.Protocol)
	}
	service.SetHandler(handler)
	if !service.Start() {
		log.Fatal(service.GetName() + " rpc service can not be start")
	}
	//RPC启动成功,则发布到中心服务器
	if !ClusterCenter.PublicService(service, config.Unique) {
		log.Fatal(service.GetName() + " rpc service can not be start")
	}
	return service
}


type ServiceCenter interface {
	GetNodeID() string
	ConnectCluster(config ClusterConfig)
	PublicService(service service.IService, unique bool) bool
	SubscribeServices(serviceNames ...string)
	GetAllService(serviceName string) []service.IService
	GetService(serviceName string, serviceID string) service.IService
	AllocService(serviceName string) service.IService

	IsConnect() bool
	Close()
}

type ClusterConfig struct {
	ID 		string   //集群中的节点id 需要保证整个集群中唯一
	Name    string     //集群名称，不用业务使用不同的集群
	Servers []string   //集群服务器列表
	Timeout uint
	LBS     string   //负载均衡策略  polling 轮询
	//CertFile string
	//KeyFile  string
	//CommonName string
}

