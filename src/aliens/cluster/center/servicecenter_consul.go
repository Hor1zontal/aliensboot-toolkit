///*******************************************************************************
// * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
// * All rights reserved.
// *
// * Contributors:
// *     aliens idea(xiamen) Corporation - initial API and implementation
// *     jialin.he <kylinh@gmail.com>
// *******************************************************************************/
package center
//
////服务中心，处理服务的调度和查询
//import (
//	"strings"
//	"github.com/hashicorp/consul/api"
//	"aliens/log"
//	"fmt"
//
//	"time"
//	"gopkg.in/mgo.v2/bson"
//)
//
//type ConsulServiceCenter struct {
//	node string
//	proxy *api.Client
//	serviceContainer map[string]*serviceCategory //服务容器 key 服务名
//
//}
//
//func DoRegistService(consul_addr string, monitor_addr string, service_name string, ip string, port int) {
//
//}
//
//
////GetNodeID() string
////ConnectCluster(config ClusterConfig)
////PublicService(service IService, unique bool) bool
////SubscribeServices(serviceNames ...string)
////GetAllService(serviceName string) []IService
////GetService(serviceName string, serviceID string) IService
////AllocService(serviceName string) IService
////
////IsConnect() bool
////Close()
//
//func (this *ConsulServiceCenter) ConnectCluster(config ClusterConfig) {
//	if config.ID == "" {
//		config.ID = bson.NewObjectId().Hex()
//	}
//	this.node = config.ID
//
//	consulConfig := api.DefaultConfig()
//	consulConfig.Address = config.Servers[0]
//
//	client, err := api.NewClient(consulConfig)
//	if err != nil {
//		log.Fatal(err)
//	}
//	this.proxy = client
//	this.serviceContainer = make(map[string]*serviceCategory)
//}
//
//func (this *ConsulServiceCenter) IsConnect() bool {
//	return this.proxy != nil
//}
//
//func (this *ConsulServiceCenter) Close() {
//	if this.proxy != nil {
//		this.proxy = nil
//	}
//}
//
//func (this *ConsulServiceCenter) PublicService(service IService, unique bool) {
//	var tags []string
//	consuleService := &api.AgentServiceRegistration{
//		ID:      this.node,
//		Name:    service.GetName(),
//		Address: service.GetAddress(),
//		Port:    service.GetPort(),
//		Tags:    tags,
//		//Check: &api.AgentServiceCheck{
//		//	HTTP:     "http://" + "monitor" + "/status",
//		//	Interval: "5s",
//		//	Timeout:  "1s",
//		//},
//	}
//	if err := this.proxy.Agent().ServiceRegister(consuleService); err != nil {
//		log.Fatal(err)
//	}
//	log.Printf("Registered service %q in consul with tags %q", service.GetName(), strings.Join(tags, ","))
//
//}
//
//func (this *ConsulServiceCenter) SubscribeServices(serviceNames ...string) {
//	t := time.NewTicker(time.Second * 5)
//	for {
//		select {
//		case <-t.C:
//			for _, serviceName := range serviceNames {
//				this.DiscoverServices( true, serviceName)
//			}
//
//		}
//	}
//}
//
//
//func (this *ConsulServiceCenter)DiscoverServices(healthyOnly bool, serviceName string) {
//	servicesData, _, err := this.proxy.Health().Service(serviceName, "", healthyOnly, &api.QueryOptions{})
//	if err == nil  {
//
//	}
//	for _, entry := range servicesData {
//		if serviceName != entry.Service.Service {
//			continue
//		}
//		for _, health := range entry.Checks {
//			if health.ServiceName != serviceName {
//				continue
//			}
//			fmt.Println("  health nodeid:", health.Node, " serviceName:", health.ServiceName, " service_id:", health.ServiceID, " status:", health.Status, " ip:", entry.Service.Address, " port:", entry.Service.Port)
//			node := newService1(health.ServiceID, serviceName, entry.Service.Address, entry.Service.Port, GRPC)
//
//
//			////get data from kv store
//			//s := GetKeyValue(serviceName, node.IP, node.Port)
//			//if len(s) > 0 {
//			//	var data KVData
//			//	err = json.Unmarshal([]byte(s), &data)
//			//	if err == nil {
//			//		node.Load = data.Load
//			//		node.Timestamp = data.Timestamp
//			//	}
//			//}
//			//fmt.Println("service node updated ip:", node.IP, " port:", node.Port, " serviceid:", node.ServiceID, " load:", node.Load, " ts:", node.Timestamp)
//			//sers = append(sers, node)
//		}
//	}
//
//	//service_locker.Lock()
//	//servics_map[serviceName] = sers
//	//service_locker.Unlock()
//}
//
//
//
////ConnectCluster(config ClusterConfig)
////PublicService(service IService, unique bool) bool
////SubscribeService(name string)
////GetAllService(name string) []IService
////IsConnect() bool
////Close()
