/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/3/28
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package service

import (
	"time"
	"aliens/module/gate/conf"
	"net"
	"aliens/module/gate/route"
	"aliens/log"
	"aliens/protocol/base"
	"aliens/common/util"
	"aliens/gate"
	"errors"
	"fmt"
)

func NewNetwork(agent gate.Agent) *Network {
	network := &Network{agent: agent, createTime:time.Now(), heartbeatTime:time.Now()}
	network.hashKey = agent.RemoteAddr().String()
	network.bindRoutes = make(map[uint16]string)
	return network
}

type Network struct {
	agent 	      gate.Agent
	//channel       chan *base.Any //消息管道

	authID  int64  //用户标识 登录验证后
	hashKey string //用来做一致性负载均衡的标识

	createTime    time.Time //创建时间
	heartbeatTime time.Time //上次的心跳时间

	bindRoutes map[uint16]string //绑定路由表 对应服务消息转发到指定节点上 比如场景服务器需要固定转发服务器
}


type IAuthMessage interface {
	GetUserID() uint32
}

//发送消息给客户端
func (this *Network) SendMessage(msg interface{}) {
	this.agent.WriteMsg(msg)
}

func (this *Network) AcceptMessage(msg *base.Any) {
	response := this.handleMessage(msg)
	if response != nil {
		this.agent.WriteMsg(response)
	}
	//this.channel <- msg
}

//绑定服务节点,固定转发
func (this *Network) BindServiceNode(serviceName string, serviceNode string) error {
	serviceSeq := route.GetServiceSeq(serviceName)
	if serviceSeq <= 0 {
		return errors.New(fmt.Sprintf("bind service node error , service %v seq not found", serviceName))
	}
	this.bindRoutes[serviceSeq] = serviceNode
	return nil
}

func (this *Network) handleMessage(request *base.Any) *base.Any {
	//未授权之前需要传递验权id
	if this.IsAuth() {
		request.AuthId = this.authID
	} else {
		request.AuthId = 0
	}

	node, ok := this.bindRoutes[request.Id]
	var response *base.Any = nil
	var error error = nil
	if ok {
		response, error = route.HandleNodeMessage(request, node)
	} else {
		response, error = route.HandleMessage(request, this.hashKey)
	}
	//log.Debugf("request %v - response %v", request, response)
	if error != nil {
		//TODO 返回服务不可用等处理方式
		log.Debug(error.Error())
	}
	//更新验权id
	if response.GetAuthId() > 0 {
		this.Auth(response.GetAuthId())
	}
	return response
}

func (this *Network) GetRemoteAddr() net.Addr {
	return this.agent.RemoteAddr()
}

func (this *Network) IsAuth() bool {
	return this.authID != 0
}

func (this *Network) Auth(id int64) {
	this.authID = id
	this.hashKey = util.Int64ToString(id)
	Manager.auth(id, this)
	//Skeleton.ChanRPCServer.Go(CommandAgentAuth, id, this)
}

//是否没有验权超时 释放多余的空连接
func (this *Network) IsAuthTimeout() bool {
	return !this.IsAuth() && time.Now().Sub(this.createTime).Seconds() >= conf.Config.AuthTimeout
}

//是否心跳超时
func (this *Network) IsHeartbeatTimeout() bool {
	return time.Now().Sub(this.heartbeatTime).Seconds() >= conf.Config.HeartbeatTimeout
}

func (this *Network) HeartBeat () {
	this.heartbeatTime = time.Now()
}