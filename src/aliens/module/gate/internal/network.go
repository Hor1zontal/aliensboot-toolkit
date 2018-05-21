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
	"aliens/common/util"
	"aliens/cluster/center"
)

var id int64 = 0

func genClientID() string {
	id ++
	return center.ClusterCenter.GetNodeID() + "_" + util.Int64ToString(id)
}

func newNetwork(agent gate.Agent) *network {
	network := &network{id: genClientID(), agent: agent, createTime:time.Now(), heartbeatTime:time.Now()}
	//network.ChannelMessageHandler = message.OpenChannelHandler(outerChannel, network, conf.Config.MessageChannelLimit)
	return network
}

type network struct {
	//*message.ChannelMessageHandler
	agent 	      gate.Agent
	id 			  string    //clientID 网络连接句柄id
	authID        int32     //验证通过的用户id 没有验证通过为0
	createTime    time.Time //创建时间
	heartbeatTime time.Time //上次的心跳时间
}

type IAuthMessage interface {
	GetUserID() uint32
}

func (this *network) GetID() string {
	return this.id
}

//发送消息给客户端
func (this *network) SendMessage(msg interface{}) {
	this.agent.WriteMsg(msg)
}

func (this *network) AcceptMessage(msg interface{}) {
	response := this.HandleMessage(msg)
	if response != nil {
		this.agent.WriteMsg(response)
	}
}

func (this *network) HandleMessage(request interface{}) interface{} {
	clientID := ""
	//未授权之前需要传递连接句柄编号
	if !this.IsAuth() {
		clientID = this.id
	}
	response, authID, error := route.HandleMessage(request, clientID)
	//TODO 返回服务不可用等处理方式
	if error != nil {
		log.Debug(error.Error())
	}

	if authID != 0 {
		this.Auth(authID)
	}
	return response
}

func (this *network) GetRemoteAddr() net.Addr {
	return this.agent.RemoteAddr()
}

func (this *network) IsAuth() bool {
	return this.authID != 0
}

func (this *network) Auth(id int32) {
	this.authID = id
	Skeleton.ChanRPCServer.Go(CommandAgentAuth, this)
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


