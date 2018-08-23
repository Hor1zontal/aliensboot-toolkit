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
	"aliens/module/gate/route"
	"aliens/log"
	"aliens/protocol/base"
	"aliens/common/util"
	"aliens/gate"
	"errors"
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"net"
)

type network struct {
	agent 	      gate.Agent
	//channel       chan *base.Any //消息管道

	authID  int64  //用户标识 登录验证后
	hashKey string //用来做一致性负载均衡的标识

	createTime    time.Time //创建时间
	heartbeatTime time.Time //上次的心跳时间

	pid *actor.PID //actor 句柄

	bindRoutes map[uint16]string //绑定路由表 对应服务消息转发到指定节点上 比如场景服务器需要固定转发服务器
}


//发送消息给客户端
func (this *network) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *base.Any:
		this.acceptMessage(msg)
	case *userPush:
		this.agent.WriteMsg(msg.pushMsg)
	case *NetworkInit:
		this.agent = msg.agent
		this.createTime = time.Now()
		this.heartbeatTime = time.Now()
		this.agent = msg.agent
		this.hashKey = this.agent.RemoteAddr().String()
		this.bindRoutes = make(map[uint16]string)
		this.pid = msg.pid
	case *actor.Stopped:
		log.Debugf("%v stop", this.getRemoteAddr())
		//TODO 如果授权 通知其他模块用户下线
		chanRpc.Go(CommandAgentRemote, this.authID, this.pid)
	}
}

func (this *network) acceptMessage(msg *base.Any) {
	response := this.handleMessage(msg)
	if response != nil {
		this.agent.WriteMsg(response)
	}
	//this.channel <- msg
}

//绑定服务节点,固定转发
func (this *network) bindServiceNode(serviceName string, serviceNode string) error {
	serviceSeq := route.GetServiceSeq(serviceName)
	if serviceSeq <= 0 {
		return errors.New(fmt.Sprintf("bind service node error , service %v seq not found", serviceName))
	}
	this.bindRoutes[serviceSeq] = serviceNode
	return nil
}

func (this *network) handleMessage(request *base.Any) *base.Any {
	//未授权之前需要传递验权id
	if this.isAuth() {
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
		this.auth(response.GetAuthId())
	}
	return response
}

func (this *network) getRemoteAddr() net.Addr {
	return this.agent.RemoteAddr()
}

func (this *network) isAuth() bool {
	return this.authID != 0
}

func (this *network) auth(id int64) {
	this.authID = id
	this.hashKey = util.Int64ToString(id)
	chanRpc.Go(CommandAgentAuth, this.authID, this.pid)
	//Skeleton.ChanRPCServer.Go(CommandAgentAuth, id, this)
}

//是否没有验权超时 释放多余的空连接
func (this *network) IsAuthTimeout() bool {
	return !this.isAuth() && time.Now().Sub(this.createTime).Seconds() >= conf.Config.AuthTimeout
}

//是否心跳超时
func (this *network) IsHeartbeatTimeout() bool {
	return time.Now().Sub(this.heartbeatTime).Seconds() >= conf.Config.HeartbeatTimeout
}

func (this *network) HeartBeat () {
	this.heartbeatTime = time.Now()
}