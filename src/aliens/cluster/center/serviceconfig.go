/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/4/12
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package center


type ServiceConfig struct {
	ID   string		//服务器的id
	Name string     //服务名称
	Address string  //服务地址 域名或ip
	Port int      //服务端端口
	Unique bool   //是否全局唯一
	Protocol string //提供服务的协议 GRPC HTTP WBSOCKET
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

type ClusterConfig struct {
	ID 		string   //集群中的节点id 需要保证整个集群中唯一
	Name    string     //集群名称，不用业务使用不同的集群
	Servers []string   //集群服务器列表
	Timeout uint
	LBS     string   //负载均衡策略  polling 轮询
	//CertFile string
	//KeyFile  string
	//CommonName string
}