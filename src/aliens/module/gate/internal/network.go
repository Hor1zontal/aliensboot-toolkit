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
	"reflect"
	"aliens/log"
	"aliens/module/gate/conf"
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

func (this *network) HandleMessage(request interface{}) interface{} {
	requestType := reflect.TypeOf(request)
	messageService := router[requestType]
	if messageService == nil {
		log.Debug("unexpect request : %v", request)
		//TODO 返回错误信息，或者T人
		return nil
	}
	//response := reflect.NewTimeWheel(responseType).Elem().Interface()
	response, error := messageService.HandleMessage(request)
	if error != nil {
		log.Debug("handle service error : %v", error)
		//TODO 返回错误信息，或者T人
	}
	return response
}

func (this *network) IsAuth() bool {
	return this.id != 0
}

//是否没有验权超时 释放多余的空连接
func (this *network) IsAuthTimeout() bool {
	return !this.IsAuth() && time.Now().Sub(this.createTime).Seconds() >= conf.Config.HeartbeatTimeout
}

//是否心跳超时
func (this *network) IsHeartbeatTimeout() bool {
	return time.Now().Sub(this.heartbeatTime).Seconds() >= conf.Config.HeartbeatTimeout
}

func (this *network) HeartBeat () {
	this.heartbeatTime = time.Now()
}


