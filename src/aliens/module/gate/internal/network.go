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
	"aliens/cluster/message"
	"time"
	"aliens/module/gate/conf"
	"github.com/name5566/leaf/gate"
	"net"
	"aliens/module/gate/route"
	"aliens/log"
)

func newNetwork(outerChannel message.IMessageChannel) *network {
	network := &network{createTime:time.Now(), heartbeatTime:time.Now()}
	network.ChannelMessageHandler = message.OpenChannelHandler(outerChannel, network, conf.Config.MessageChannelLimit)
	return network
}

type network struct {
	*message.ChannelMessageHandler
	id int32 //验证通过的用户id 没有验证通过为0
	createTime time.Time //创建时间
	heartbeatTime time.Time //上次的心跳时间
}

type IAuthMessage interface {
	GetID() int32
}

func (this *network) HandleMessage(request interface{}) interface{} {
	response, error := route.HandleMessage(request)
	//TODO 返回服务不可用 或 嘿嘿嘿
	if error != nil {
		log.Debug(error.Error())
	}
	return response
}

func (this *network) GetRemoteAddr() net.Addr {
	channel := this.GetOuterChannel()
	if channel == nil {
		return nil
	}
	return channel.(gate.Agent).RemoteAddr()
}

func (this *network) IsAuth() bool {
	return this.id != 0
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


