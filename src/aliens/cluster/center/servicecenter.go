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
	"aliens/cluster/center/lbs"
)

var ClusterCenter ServiceCenter = &ETCDServiceCenter{} //服务调度中心

const NODE_SPLIT string = "/"

const SERVICE_NODE_NAME string = "service"

const CONFIG_NODE_NAME string = "config"

const DEFAULT_LBS string = lbs.LbsStrategyPolling

func PublicService(config service.Config, handler interface{}) service.IService {
	if !ClusterCenter.IsConnect() {
		log.Fatal(config.Name + " cluster center is not connected")
		return nil
	}

	config.ID = ClusterCenter.GetNodeID() //节点id

	service := service.NewService(config)
	if service == nil {
		log.Fatalf( "un expect service protocol %v", config.Protocol)
	}
	service.SetHandler(handler)
	if !service.Start() {
		log.Fatalf("service %v can not be start", service.GetName())
	}
	//RPC启动成功,则发布到中心服务器
	if !ClusterCenter.PublicService(service, config.Unique) {
		log.Fatalf("service %v can not be public", service.GetName())
	}
	return service
}


func ReleaseService(service service.IService) {
	if service != nil {
		service.Close()
	}
	if !ClusterCenter.IsConnect() {
		log.Errorf(" cluster center is not connected")
		return
	}
	ClusterCenter.ReleaseService(service)
}

type ConfigListener func(data []byte)

type ServiceCenter interface {

	GetNodeID() string

	ConnectCluster(config ClusterConfig)

	PublicConfig(configName string, content []byte) bool        //发布配置
	SubscribeConfig(configName string, listener ConfigListener) //订阅配置

	ReleaseService(service service.IService)   //释放服务
	PublicService(service service.IService, unique bool) bool  //发布服务

	SubscribeServices(serviceNames ...string) //订阅服务
	GetAllService(serviceName string) []service.IService  //获取所有的服务

	GetService(serviceName string, serviceID string) service.IService //获取指定服务
	AllocService(serviceName string, param string) service.IService

	//AddServiceListener(listener service.Listener)

	IsConnect() bool
	Close()
}

type ClusterConfig struct {
	ID 		string   //集群中的节点id 需要保证整个集群中唯一
	Name    string     //集群名称，不用业务使用不同的集群
	Servers []string   //集群服务器列表
	Timeout uint
	LBS     string   //负载均衡策略  polling 轮询
	TTL     int64   //
	//CertFile string
	//KeyFile  string
	//CommonName string
}

