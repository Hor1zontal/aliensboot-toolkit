/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/6/4
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import "aliens/config"

func NewService(config config.ServiceConfig) IService {
	return NewService1(config.ID, config.Name, config.Address, config.Port, config.Protocol)
}

func NewService2(centerService *CenterService, id string, name string) IService {
	centerService.SetID(id)
	centerService.SetName(name)
	var service IService = nil
	//switch centerService.Protocol {
	//case GRPC:
	//	service = &GRPCService{CenterService: centerService}
	//default:
	//
	//	//case WEBSOCKET:
	//	//	return &wbService{centerService: centerService}
	//	//case HTTP:
	//	//	return &httpService{centerService: centerService}
	//}
	service = &GRPCService{CenterService: centerService}
	return service
}

func NewService1(id string, name string, address string, port int, protocol string) IService {
	centerService := &CenterService{
		Address:  address,
		Port:     port,
		Protocol: protocol,
	}
	return NewService2(centerService, id, name)
}