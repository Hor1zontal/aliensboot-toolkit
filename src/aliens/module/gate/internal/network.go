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
	"aliens/protocol"
)

var id int64 = 0

//func genClientID() string {
//	id ++
//	return center.ClusterCenter.GetNodeID() + "_" + util.Int64ToString(id)
//}

func newNetwork(agent gate.Agent) *network {
	network := &network{agent: agent, createTime:time.Now(), heartbeatTime:time.Now()}
	network.channel = make(chan *CallInfo, 5)
	go func() {
		for {
			info, ok := <-network.channel
			if !ok {
				return
			}
			response := network.HandleMessage(info.msg)
			if response != nil {
				info.agent.WriteMsg(response)
			}
		}
	}()
	return network
}

type network struct {
	agent 	      gate.Agent
	channel       chan *CallInfo //消息管道

	authID        int64     //验证通过的用户id 没有验证通过为0
	createTime    time.Time //创建时间
	heartbeatTime time.Time //上次的心跳时间
}

type CallInfo struct {
	msg   *protocol.Any
	agent gate.Agent
}

type IAuthMessage interface {
	GetUserID() uint32
}

//发送消息给客户端
func (this *network) SendMessage(msg interface{}) {
	this.agent.WriteMsg(msg)
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

func (this *network) AcceptMessage(msg *protocol.Any) {
	this.channel <- &CallInfo{msg, this.agent}
}

func (this *network) HandleMessage(request *protocol.Any) *protocol.Any {
	//未授权之前需要传递连接句柄编号
	if !this.IsAuth() {

	} else {
		request.AuthId = this.authID
	}
	response, error := route.HandleMessage(request)
	if error != nil {
		//TODO 返回服务不可用等处理方式
		log.Debug(error.Error())
	}
	if response.GetAuthId() != 0 {
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


