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
	"aliens/log"
	"encoding/json"
)

func newCenterService(config ServiceConfig) *centerService {
	if !ClusterCenter.IsConnect() {
		log.Fatal(config.Name + " cluster center is not connected")
		return nil
	}
	return &centerService{
		id:       ClusterCenter.GetNodeID(),
		name:     config.Name,
		Ip:  	  config.Address,
		Port:     config.Port,
		Protocol: config.Protocol,
	}
}

func loadCenterService(data []byte, id string, name string) IService {
	centerService := &centerService{}
	json.Unmarshal(data, centerService)
	centerService.SetID(id)
	centerService.SetName(name)

	switch centerService.Protocol {
	case GRPC:
		return &GRPCService{centerService: centerService}
		//case WEBSOCKET:
		//	return &wbService{centerService: centerService}
		//case HTTP:
		//	return &httpService{centerService: centerService}
	}
	return nil
}

//func newService(service *centerService) IService {
//	switch service.Protocol {
//		case GRPC:
//		return &GRPCService{centerService: service}
//		//case WEBSOCKET:
//		//	return &wbService{centerService: centerService}
//		//case HTTP:
//		//	return &httpService{centerService: centerService}
//		}
//	return nil
//}

func PublicService(config ServiceConfig, handler interface{}) IService {
	proxy := newCenterService(config)
	var service IService = nil
	switch config.Protocol {
	case GRPC:
		service = &GRPCService{centerService: proxy}
		break
		//case WEBSOCKET:
		//	return &wbService{centerService: centerService}
		//case HTTP:
		//	return &httpService{centerService: centerService}
	}
	if service == nil {
		log.Fatalf( "un expect service protocol %v", config.Protocol)
	}
	service.SetHandler(handler)
	if !service.Start() {
		log.Fatal(service.GetProxy().GetName() + " rpc service can not be start")
	}
	//RPC启动成功,则发布到中心服务器
	if !ClusterCenter.PublicService(service, config.Unique) {
		log.Fatal(service.GetProxy().GetName() + " rpc service can not be start")
	}
	return service
}

