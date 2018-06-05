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



func loadService(data []byte, id string, name string) IService {
	centerService := &ServiceConfig{}
	json.Unmarshal(data, centerService)
	centerService.SetID(id)
	centerService.SetName(name)

	switch centerService.Protocol {
	case GRPC:
		return &GRPCService{ServiceConfig: centerService}
		//case WEBSOCKET:
		//	return &wbService{centerService: centerService}
		//case HTTP:
		//	return &httpService{centerService: centerService}
	}
	return nil
}

func newService(config ServiceConfig) IService {
	if !ClusterCenter.IsConnect() {
		log.Fatal(config.Name + " cluster center is not connected")
		return nil
	}
	return loadService1(ClusterCenter.GetNodeID(), config.Name, config.Address, config.Port, config.Protocol)
}

func loadService1(id string, name string, address string, port int, protocol string) IService {
	centerService := &ServiceConfig{
		ID:       id,
		Name:     name,
		Address:  address,
		Port:     port,
		Protocol: protocol,
	}

	var service IService = nil
	switch centerService.Protocol {
	case GRPC:
		service = &GRPCService{ServiceConfig: centerService}
		//case WEBSOCKET:
		//	return &wbService{centerService: centerService}
		//case HTTP:
		//	return &httpService{centerService: centerService}
	}
	return service
}

func PublicService(config ServiceConfig, handler interface{}) IService {
	service := newService(config)
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

