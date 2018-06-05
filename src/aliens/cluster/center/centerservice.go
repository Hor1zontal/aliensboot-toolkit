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

//调用方式
const (
	WEBSOCKET string = "websocket"
	HTTP string = "http"
	GRPC string = "grpc"
)


type IService interface {
	GetID() string
	SetID(id string)
	GetName() string
	SetName(name string)

	GetDesc() string                                  //获取服务的描述信息
	GetProxy() *centerService


	Start() bool                                      //启动服务
	Connect() bool                                    //连接服务
	Close() 										  //关闭服务

	Equals(other IService) bool                       //比较服务
	IsLocal() bool                                    //是否本机服务
	Request(request interface{}) (interface{}, error) //请求服务
	SetHandler(handler interface{})  //设置处理句柄
	//Push(request interface{}) error //服务推送
}

type centerService struct {
	Ip       string `json:"ip"`
	Port     int	`json:"port"`
	//Address  string `json:"address"`  	//服务访问地址 写入到中心服务器供外部调用
	Protocol string `json:"protocol"` 		//服务的访问方式
	id       string `json:"-"`              //服务ID
	name     string `json:"-"`              //服务类型
}


func (this *centerService) SetID(id string) {
	this.id = id
}


func (this *centerService) GetID() string {
	return this.id
}



func (this *centerService) GetName() string {
	return this.name
}


func (this *centerService) SetName(name string) {
	this.name = name
}