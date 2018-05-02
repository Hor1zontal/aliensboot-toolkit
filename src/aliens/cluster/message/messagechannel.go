/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/5/6
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package message

import (
	"aliens/log"
	"aliens/common/util"
)

type IMessageChannel interface {
	WriteMsg(msg interface{})
	Close()
	SetUserData(data interface{})
	UserData() interface{}
}

type IMessageHandler interface {
	HandleMessage(msg interface{}) interface{} //处理
}

type MessageChannel struct {
	channel        chan interface{} //管道
	open           bool
	messageLimit   int
	messageHandler *ChannelMessageHandler //消息处理handler
	userdata       interface{}
}

//向管道发送消息
func (this *MessageChannel) WriteMsg(message interface{}) {
	//用户消息管道没开，不接受消息
	if this.channel == nil {
		return
	}
	select {
	case this.channel <- message:
	default:
		log.Debugf("message channel full %v - %v", this.channel, message)
		//TODO 消息管道满了需要异常处理
	}
}

//打开用户消息管道
func (this *MessageChannel) EnsureOpen() {
	if this.channel != nil {
		return
	}
	this.channel = make(chan interface{}, this.messageLimit)
	go func() {
		defer func() {
			util.CatchStackDetail()
		}()

		for {
			//只要消息管道没有关闭，就一直等待用户请求
			message, open := <-this.channel
			if !this.open || !open {
				this.channel = nil
				break
			}
			this.messageHandler.HandleMessage(message)
		}
		this.Close()
	}()
	this.open = true
}

//关闭消息管道
func (this *MessageChannel) Close() {
	if !this.IsOpen() {
		return
	}
	defer func() {
		recover()
	}()
	this.open = false
	if this.channel != nil {
		close(this.channel)
	}
}

//消息管道是否打开
func (this *MessageChannel) IsOpen() bool {
	return this.open
}

func (this *MessageChannel) SetUserData(data interface{}) {
	this.userdata = data
}

func (this *MessageChannel) UserData() interface{} {
	return this.userdata
}
