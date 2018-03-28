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
	GetDesc() string //获取服务的描述信息
	GetID() string   //获取id
	SetID(id string) //设置id
	GetType() string //获取服务类型
	SetType(serviceType string) //设置服务类型
	Start() bool                              //启动服务
	Connect() bool                            //连接服务
	Equals(other IService) bool               //比较服务
	IsLocal() bool                            //是否本机服务
	Request(request interface{}) (interface{}, error) //服务请求
	//Push(request interface{}) error //服务推送
}

type centerService struct {
	Address     string `json:"address"` //服务访问地址 写入到中心服务器供外部调用
	Protocol    string `json:"protocol"` //服务的访问方式
	id          string //服务ID
	serviceType string //服务类型
}