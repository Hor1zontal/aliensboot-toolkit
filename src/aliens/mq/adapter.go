/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/4/8
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package mq

import "aliens/mq/kafka"

type Type uint8

const (
	TYPE_KAFKA Type = 1
)

//消息生产者
type IProducer interface {
	Init(address []string, timeout int) error
	SendMessage(service string, node string, data []byte) //异步发送数据
	Broadcast(service string, data []byte) //广播数据
	Close() error
}

//消息消费者
type IConsumer interface {
	Init(address []string, service string, node string, handle func(data []byte) error) error
	Close() error
}

func NewProducer(mqType Type, config Config) (producer IProducer, err error) {
	if mqType == TYPE_KAFKA {
		producer = &kafka.Producer{}
	}
	if producer != nil {
		err = producer.Init(config.Address, config.Timeout)
	}
	return
}

func NewConsumer(mqType Type, config Config, service string, node string, handle func(data []byte) error) (consumer IConsumer, err error) {
	if mqType == TYPE_KAFKA {
		consumer = &kafka.Consumer{}
	}

	if consumer != nil {
		err = consumer.Init(config.Address, service, node, handle)
	}
	return
}