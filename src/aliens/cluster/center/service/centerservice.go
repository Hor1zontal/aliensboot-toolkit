/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/3/24
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

//调用方式
const (
	WEBSOCKET string = "websocket"
	HTTP string = "http"
	GRPC string = "grpc"
)


type IService interface {
	GetID() string
	SetID(id string)
	SetName(name string)
	GetName() string
	GetAddress() string
	GetPort() int
	//GetConfig() *ServiceConfig


	Start() bool                                      //启动服务
	Connect() bool                                    //连接服务
	Close() 										  //关闭服务

	Equals(other IService) bool                       //比较服务
	IsLocal() bool                                    //是否本机服务
	Request(request interface{}) (interface{}, error) //请求服务
	SetHandler(handler interface{})  //设置处理句柄
	//Push(request interface{}) error //服务推送
}


type ServiceConfig struct {
	ID   string		//服务器的id
	Name string     //服务名称
	Address string  //服务地址 域名或ip
	Port int      //服务端端口
	Unique bool   //是否全局唯一
	Protocol string //提供服务的协议 GRPC HTTP WBSOCKET
}


func (this *ServiceConfig) GetAddress() string {
	return this.Address
}


func (this *ServiceConfig) SetAddress(address string) {
	this.Address = address
}

func (this *ServiceConfig) GetPort() int {
	return this.Port
}

func (this *ServiceConfig) SetPort(port int) {
	this.Port = port
}


func (this *ServiceConfig) GetID() string {
	return this.ID
}


func (this *ServiceConfig) SetID(id string) {
	this.ID = id
}


func (this *ServiceConfig) GetName() string {
	return this.Name
}


func (this *ServiceConfig) SetName(name string) {
	this.Name = name
}


//
//type centerService struct {
//	Ip       string `json:"ip"`
//	Port     int	`json:"port"`
//	//Address  string `json:"address"`  	//服务访问地址 写入到中心服务器供外部调用
//	Protocol string `json:"protocol"` 		//服务的访问方式
//	id       string `json:"-"`              //服务ID
//	name     string `json:"-"`              //服务类型
//}


