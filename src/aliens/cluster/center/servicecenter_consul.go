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
//	"encoding/json"
//	"time"
//	"gopkg.in/mgo.v2/bson"
//)
//
//type ConsulServiceCenter struct {
//	proxy *api.Client
//	node string
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
//
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
//	serviceName := service.GetName()
//
//	var tags []string
//	consuleService := &api.AgentServiceRegistration{
//		ID:      this.node,
//		Name:    serviceName,
//		Address: service.GetConfig().Address,
//		Port:    service.GetConfig().Port,
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
//	log.Printf("Registered service %q in consul with tags %q", serviceName, strings.Join(tags, ","))
//
//}
//
//func (this *ConsulServiceCenter) SubscribeServices(serviceNames ...string) {
//	t := time.NewTicker(time.Second * 5)
//	for {
//		select {
//		case <-t.C:
//			this.DiscoverServices( true, found_service)
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
//
//			var node IService
//			node.SetID(health.ServiceID)
//			node.SetName(serviceName)
//			node.Start()
//
//			node.IP = entry.Service.Address
//			node.Port = entry.Service.Port
//			node.ServiceID = health.ServiceID
//
//			//get data from kv store
//			s := GetKeyValue(serviceName, node.IP, node.Port)
//			if len(s) > 0 {
//				var data KVData
//				err = json.Unmarshal([]byte(s), &data)
//				if err == nil {
//					node.Load = data.Load
//					node.Timestamp = data.Timestamp
//				}
//			}
//			fmt.Println("service node updated ip:", node.IP, " port:", node.Port, " serviceid:", node.ServiceID, " load:", node.Load, " ts:", node.Timestamp)
//			sers = append(sers, node)
//		}
//	}
//
//	service_locker.Lock()
//	servics_map[serviceName] = sers
//	service_locker.Unlock()
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
