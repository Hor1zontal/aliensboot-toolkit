/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/3/24
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package center


func PublicWBService(serviceType string, address string) *wbService {
	if !ClusterCenter.IsConnect() {
		panic(serviceType + " cluster center is not connected")
		return nil
	}
	service := &wbService{
		&centerService{
			id:          ClusterCenter.GetNodeID(),
			serviceType: serviceType,
			Address:     address,
			Protocol: WEBSOCKET,
		},
	}
	//center.ClusterCenter.AddServiceFactory(service.serviceType, &wbServiceFactory{})
	//websocket服务启动成功,则发布到中心服务器
	if !ClusterCenter.PublicService(service) {
		panic(service.serviceType + " wb service can not be public")
	}
	return service
}

type wbService struct {
	*centerService
}

func (this *wbService) GetDesc() string {
	return "websocket service"
}

func (this *wbService) GetID() string {
	return this.id
}

func (this *wbService) GetType() string {
	return this.serviceType
}

func (this *wbService) SetID(id string) {
	this.id = id
}

func (this *wbService) SetType(serviceType string) {
	this.serviceType = serviceType
}

//启动服务
func (this *wbService) Start() bool {
	return true
}

//连接服务
func (this *wbService) Connect() bool {
	return true
}

//比较服务是否冲突
func (this *wbService) Equals(other IService) bool {
	otherService, ok := other.(*wbService)
	if !ok {
		return false
	}
	return this.serviceType == otherService.serviceType && this.Address == otherService.Address
}

//服务是否本进程启动的
func (this *wbService) IsLocal() bool {
	return true
}

//关闭服务
func (this *wbService) Close() {
}

//向服务请求消息
func (this *wbService) Request(request interface{}) (interface{}, error) {
	return nil, nil
}
