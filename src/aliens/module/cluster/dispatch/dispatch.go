/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/4/23
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package dispatch

import (
	"aliens/protocol"
	"fmt"
	"aliens/mq"
	"aliens/module/cluster/conf"
	"aliens/log"
	"github.com/gogo/protobuf/proto"
	"aliens/module/cluster/cache"
	"errors"
	//"aliens/module/cluster"
	"aliens/cluster/center"
	"aliens/module/cluster/constant"
)

//消息生产者
var mqProducer mq.IProducer = nil

//消息消费者
var mqConsumer = make(map[string]mq.IConsumer)


func Init() {
	center.ClusterCenter.ConnectCluster(conf.Config.Cluster)
	handler, err := mq.NewProducer(mq.TYPE_KAFKA, conf.Config.MQ)
	if err != nil {
		log.Fatal("%v", err)
	} else {
		mqProducer = handler
	}
	log.Info("init message producer success")
}

func Close() {
	center.ClusterCenter.Close()
}

//注册消息队列消费者 一般用来处理推送消息
func RegisterConsumer(serviceType string, handle func(data *protocol.Any) error) {
	consumerID := serviceType + center.ClusterCenter.GetNodeID()
	consumer := mqConsumer[consumerID]
	if consumer != nil {
		log.Warn("consumer %v already register", consumerID)
		return
	}

	handleProxy := NewProtobufHandler(handle).HandleMessage
	consumer, err := mq.NewConsumer(mq.TYPE_KAFKA, conf.Config.MQ, serviceType, center.ClusterCenter.GetNodeID(), handleProxy)
	if err != nil {
		log.Fatal("%v", err)
	} else {
		mqConsumer[consumerID] = consumer
		log.Info("register consumer %v success", consumerID)
	}
}

//注销消费者
func UnregisterConsumer(serviceType string) {
	consumerID := serviceType + center.ClusterCenter.GetNodeID()
	consumer := mqConsumer[consumerID]
	if consumer != nil {
		err := consumer.Close()
		if err != nil {
			log.Error("%v", err)
		}
	}
}

//网关异步推送信息给指定用户
func GatePush(serviceType string, clientID string, message proto.Message) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	request := &protocol.Any{TypeUrl: serviceType, SessionId: clientID, Value: data}
	gateID := cache.ClusterCache.GetClientGateID(clientID)
	if gateID == "" {
		return errors.New(fmt.Sprint("gate ID can not found, clientID : %v", clientID))
	}
	return AsyncPush(constant.SERVICE_GATE, gateID, request)
}

//消息异步推送 - 推送到指定服务节点
func AsyncPush(serviceType string, serviceID string, message proto.Message) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	mqProducer.SendMessage(serviceType, serviceID, data)
	return nil
}

//消息异步推送 - 广播所有对应服务节点
func AsyncBroadcast(serviceType string, message proto.Message) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	mqProducer.Broadcast(serviceType, data)
	return nil
}

