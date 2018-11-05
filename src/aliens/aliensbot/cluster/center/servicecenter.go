/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/6/4
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package center

import (
	"aliens/aliensbot/cluster/center/lbs"
	"aliens/aliensbot/cluster/center/service"
	"aliens/aliensbot/common/util"
	"aliens/aliensbot/config"
	"aliens/aliensbot/log"
)

var ClusterCenter ServiceCenter = &ETCDServiceCenter{} //服务调度中心

const NODE_SPLIT string = "/"

const SERVICE_NODE_NAME string = "service"

const CONFIG_NODE_NAME string = "config"

const DEFAULT_LBS string = lbs.LbsStrategyPolling

func PublicService(config config.ServiceConfig, handler interface{}) service.IService {
	if !ClusterCenter.IsConnect() {
		log.Fatal(config.Name + " cluster center is not connected")
		return nil
	}
	config.ID = ClusterCenter.GetNodeID() //节点id
	//地址没有发布到外网 采用内网地址
	if config.Address == "" {
		config.Address = util.GetIP()
	}
	service := service.NewService(config)
	if service == nil {
		log.Fatalf("un expect service protocol %v", config.Protocol)
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
	//if !ClusterCenter.IsConnect() {
	//	log.Errorf(" cluster center is not connected")
	//	return
	//}
	//先从中心释放，再内部关闭，缓解关闭期间其他服务请求转发过来
	ClusterCenter.ReleaseService(service)
	if service != nil {
		service.Close()
	}
}

type ConfigListener func(data []byte)

type ServiceCenter interface {
	GetNodeID() string //获取当前节点id

	ConnectCluster(config config.ClusterConfig)

	PublicConfig(configName string, content []byte) bool        //发布配置
	SubscribeConfig(configName string, listener ConfigListener) //订阅配置

	ReleaseService(service service.IService)                  //释放服务
	PublicService(service service.IService, unique bool) bool //发布服务

	SubscribeServices(serviceNames ...string)            //订阅服务
	GetAllService(serviceName string) []service.IService //获取所有的服务

	GetService(serviceName string, serviceID string) service.IService //获取指定服务
	AllocService(serviceName string, param string) service.IService   //按照负载均衡策略 分配一个可用的服务

	//AddServiceListener(listener service.Listener)

	IsConnect() bool
	Close()
}