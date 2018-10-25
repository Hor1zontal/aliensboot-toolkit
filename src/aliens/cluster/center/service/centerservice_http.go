/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/3/24
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service


//func PublicHTTPService(config CenterService, address string) *httpService {
//	if !ClusterCenter.IsConnect() {
//		panic(config.Name + " cluster center is not connected")
//		return nil
//	}
//	service := &httpService{
//		&centerService{
//			id:       ClusterCenter.GetNodeID(),
//			name:     config.Name,
//			Ip:  address,
//			Port: port,
//			Protocol: HTTP,
//		},
//	}
//	//center.ClusterCenter.AddServiceFactory(service.name, &HTTPServiceFactory{})
//	//websocket服务启动成功,则发布到中心服务器
//	if !ClusterCenter.PublicService(service, config.Unique) {
//		panic(service.name + " http service can not be public")
//	}
//	return service
//}
//
//type httpService struct {
//	*centerService
//}
//
//func (this *httpService) GetDesc() string {
//	return "http service"
//}
//
//func (this *httpService) GetID() string {
//	return this.id
//}
//
//func (this *httpService) GetType() string {
//	return this.name
//}
//
//func (this *httpService) SetID(id string) {
//	this.id = id
//}
//
//func (this *httpService) SetType(serviceType string) {
//	this.name = serviceType
//}
//
////启动服务
//func (this *httpService) Start() bool {
//	return true
//}
//
////连接服务
//func (this *httpService) Connect() bool {
//	return true
//}
//
////比较服务是否冲突
//func (this *httpService) Equals(other IService) bool {
//	otherService, ok := other.(*httpService)
//	if !ok {
//		return false
//	}
//	return this.name == otherService.name && this.Ip == otherService.Ip && this.Port == otherService.Port
//}
//
////服务是否本进程启动的
//func (this *httpService) IsLocal() bool {
//	return true
//}
//
////关闭服务
//func (this *httpService) Close() {
//}
//
////向服务请求消息
//func (this *httpService) Request(request interface{}) (interface{}, error) {
//	return nil, nil
//}
