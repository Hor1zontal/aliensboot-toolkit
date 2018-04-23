/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/4/21
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package kafka

import (
	"time"
	"github.com/Shopify/sarama"
	"aliens/log"
)

type Producer struct {
	proxy sarama.AsyncProducer
}

func (this *Producer) Init(address []string, timeout int) error {
	if timeout == 0 {
		timeout = 5
	}
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true //必须有这个选项
	config.Producer.Timeout = time.Duration(timeout) * time.Second
	p, err := sarama.NewAsyncProducer(address, config)
	//defer p.Close()
	if err != nil {
		return err
	}

	this.proxy = p
	//必须有这个匿名函数内容
	go func(p sarama.AsyncProducer) {
		errors := p.Errors()
		success := p.Successes()
		for {
			select {
			case err := <-errors:
				if err != nil {
					log.Error(err)
				}
			case succ := <-success:
				if succ != nil {
					log.Debug(succ)
				}
			}
		}
	}(this.proxy)
	return nil
}

func (this *Producer) Close() error {
	if this.proxy != nil {
		return this.proxy.Close()
	}
	return nil
}

func (this *Producer) Broadcast(service string, data []byte) {
	msg := &sarama.ProducerMessage{
		Topic: service,
		Value: sarama.ByteEncoder(data),
	}
	this.proxy.Input() <- msg
}

func (this *Producer) SendMessage(service string, node string, data []byte) {
	msg := &sarama.ProducerMessage{
		Topic: service + node,
		Value: sarama.ByteEncoder(data),
	}
	this.proxy.Input() <- msg
}

