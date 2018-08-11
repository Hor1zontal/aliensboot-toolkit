/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/3/28
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package internal

import (
	"time"
	"aliens/module/gate/conf"
	"github.com/name5566/leaf/gate"
	"net"
	"aliens/module/gate/route"
	"aliens/log"
	"aliens/protocol/base"
)

//var id int64 = 0

//func genClientID() string {
//	id ++
//	return center.ClusterCenter.GetNodeID() + "_" + util.Int64ToString(id)
//}

func newNetwork(agent gate.Agent) *network {
	network := &network{agent: agent, createTime:time.Now(), heartbeatTime:time.Now()}
	network.channel = make(chan *base.Any, 5)
	go network.Start()
	return network
}

type network struct {
	agent 	      gate.Agent
	channel       chan *base.Any //消息管道

	authID        int64     //用户标识 登录验证后
	createTime    time.Time //创建时间
	heartbeatTime time.Time //上次的心跳时间

	routes map[string]string //路由表 消息服务-服务节点

	userData interface{}
}


type IAuthMessage interface {
	GetUserID() uint32
}

//发送消息给客户端
//func (this *network) SendMessage(msg interface{}) {
//	this.agent.WriteMsg(msg)
//}

func (this *network) Start() {
	for {
		msg, ok := <-this.channel
		if !ok {
			return
		}
		//msg.Agent = this
		response := this.HandleMessage(msg)
		if response != nil {
			this.agent.WriteMsg(response)
		}
	}
}

//发送消息给客户端
func (this *network) Close() (isClosed bool) {
	defer func() {
		if recover() != nil {
			isClosed = false
		}
	} ()
	close(this.channel)
	return true
}

func (this *network) AcceptMessage(msg *base.Any) {
	this.channel <- msg
}

func (this *network) HandleMessage(request *base.Any) *base.Any {
	//未授权之前需要传递验权id
	if this.IsAuth() {
		request.AuthId = this.authID
	} else {
		request.AuthId = 0
	}
	response, error := route.HandleMessage(request)
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

func (this *network) GetRemoteAddr() net.Addr {
	return this.agent.RemoteAddr()
}

func (this *network) IsAuth() bool {
	return this.authID != 0
}

func (this *network) Auth(id int64) {
	this.authID = id
	Skeleton.ChanRPCServer.Go(CommandAgentAuth, id, this)
}

//是否没有验权超时 释放多余的空连接
func (this *network) IsAuthTimeout() bool {
	return !this.IsAuth() && time.Now().Sub(this.createTime).Seconds() >= conf.Config.AuthTimeout
}

//是否心跳超时
func (this *network) IsHeartbeatTimeout() bool {
	return time.Now().Sub(this.heartbeatTime).Seconds() >= conf.Config.HeartbeatTimeout
}

func (this *network) HeartBeat () {
	this.heartbeatTime = time.Now()
}

func (this *network) WriteMsg(msg interface{}) {
	this.agent.WriteMsg(msg)
}

func (this *network) UserData() interface{} {
	return this.userData
}

func (this *network) SetUserData(data interface{}) {
	this.userData = data
}