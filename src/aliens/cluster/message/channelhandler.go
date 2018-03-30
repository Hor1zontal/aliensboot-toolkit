/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2017/5/12
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package message

import (
	"aliens/common/util"
)

func OpenChannelHandler(outerChannel IMessageChannel, handler IMessageHandler, maxMessage int) *ChannelMessageHandler {
	channelHandler := &ChannelMessageHandler{
		outerChannel: outerChannel,
		handler:      handler,
	}
	channelHandler.InitInnerChannel(maxMessage)
	return channelHandler
}

type IChannelMessageHandler interface {
	AcceptMessage(request interface{})
}

type ChannelMessageHandler struct {
	outerChannel IMessageChannel //往外写的消息管道
	innerChannel *MessageChannel //往内写的消息管道
	handler      IMessageHandler //服务处理类
}

//收取系统消息
func (this *ChannelMessageHandler) SetHandler(handler IMessageHandler) {
	this.handler = handler
}

func (this *ChannelMessageHandler) GetOuterChannel() IMessageChannel {
	return this.outerChannel
}

func (this *ChannelMessageHandler) SetOuterChannel(outerChannel IMessageChannel) {
	this.outerChannel = outerChannel
}

//func (this *ChannelMessageHandler) GateClose(gate gate.Agent) {
//	this.Close()
//}

//收取消息
func (this *ChannelMessageHandler) AcceptMessage(message interface{}) {
	if this.innerChannel != nil {
		this.innerChannel.EnsureOpen()
		this.innerChannel.WriteMsg(message)
	}
}

//往连接客户端写消息
func (this *ChannelMessageHandler) SendMessage(message interface{}) {
	if this.outerChannel != nil {
		this.outerChannel.WriteMsg(message)
	}
}

//是否在线
func (this *ChannelMessageHandler) IsOnline() bool {
	return this.outerChannel != nil
}

func (this *ChannelMessageHandler) HandleMessage(msg interface{}) {
	defer func() {
		//处理消息异常
		if err := recover(); err != nil {
			util.PrintStackDetail()
			this.Close()
		}
	}()
	if this.handler != nil {
		response := this.handler.HandleMessage(msg)
		if this.outerChannel != nil && response != nil {
			this.outerChannel.WriteMsg(response)
		}
	}
}

//打开收消息管道
func (this *ChannelMessageHandler) InitInnerChannel(maxMessage int) {
	if this.innerChannel != nil {
		return
	}
	channel := &MessageChannel{
		messageLimit:   maxMessage,
		messageHandler: this,
	}
	this.innerChannel = channel
}

//关闭收消息管道
func (this *ChannelMessageHandler) CloseInnerChannel() {
	if this.innerChannel != nil {
		this.innerChannel.Close()
		this.innerChannel = nil
	}
}

//关闭收消息管道
func (this *ChannelMessageHandler) CloseOuterChannel() {
	if this.outerChannel != nil {
		this.outerChannel.Close()
		this.outerChannel = nil
	}
}

//关闭所有管道
func (this *ChannelMessageHandler) Close() {
	this.CloseInnerChannel()
	this.CloseOuterChannel()
}
