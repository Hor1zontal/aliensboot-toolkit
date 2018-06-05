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
//)
//
//type ConsulServiceCenter struct {
//	proxy *api.Client
//}
//
//func DoRegistService(consul_addr string, monitor_addr string, service_name string, ip string, port int) {
//
//}
//
//func (this *ConsulServiceCenter) ConnectCluster(config ClusterConfig) {
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
//func (this *ConsulServiceCenter) PublicService(service IService, unique bool) {
//	serviceName := service.GetName()
//	serviceID := serviceName + "-" + service.GetIP() + ":"
//	var tags []string
//	consuleService := &api.AgentServiceRegistration{
//		ID:      serviceID,
//		Name:    serviceName,
//		Address: service.GetIP(),
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
//	log.Printf("Registered service %q in consul with tags %q", serviceName, strings.Join(tags, ","))
//
//}
//
//func (this *ConsulServiceCenter) SubscribeService(serviceTypes ...string) {
//	t := time.NewTicker(time.Second * 5)
//	for {
//		select {
//		case <-t.C:
//			DiscoverServices( true, found_service)
//		}
//	}
//}
//
//
//func (this *ConsulServiceCenter)DiscoverServices(healthyOnly bool, service_name string) {
//	services, _, err := this.proxy.Catalog().Services(&api.QueryOptions{})
//	if err != nil {
//		log.Errorf("subscribe service %v error: %v", service_name, err)
//		return
//	}
//
//	//var sers ServiceList
//	for name := range services {
//		servicesData, _, err := this.proxy.Health().Service(name, "", healthyOnly,
//			&api.QueryOptions{})
//		CheckErr(err)
//		for _, entry := range servicesData {
//			if service_name != entry.Service.Service {
//				continue
//			}
//			for _, health := range entry.Checks {
//				if health.ServiceName != service_name {
//					continue
//				}
//				fmt.Println("  health nodeid:", health.Node, " service_name:", health.ServiceName, " service_id:", health.ServiceID, " status:", health.Status, " ip:", entry.Service.Address, " port:", entry.Service.Port)
//
//				var node IService
//				node.SetID(health.ServiceID)
//				node.SetName(service_name)
//				node.Start()
//
//				node.IP = entry.Service.Address
//				node.Port = entry.Service.Port
//				node.ServiceID = health.ServiceID
//
//				//get data from kv store
//				s := GetKeyValue(service_name, node.IP, node.Port)
//				if len(s) > 0 {
//					var data KVData
//					err = json.Unmarshal([]byte(s), &data)
//					if err == nil {
//						node.Load = data.Load
//						node.Timestamp = data.Timestamp
//					}
//				}
//				fmt.Println("service node updated ip:", node.IP, " port:", node.Port, " serviceid:", node.ServiceID, " load:", node.Load, " ts:", node.Timestamp)
//				sers = append(sers, node)
//			}
//		}
//	}
//
//	service_locker.Lock()
//	servics_map[service_name] = sers
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
