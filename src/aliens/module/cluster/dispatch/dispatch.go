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
	"aliens/module/cluster/conf"
	"aliens/cluster/center"
)

var RPC = newGRPCDispatcher()

var MQ = newMQDispatcher(conf.Config.MQ)

//type Dispatcher interface {
//	//RegisterConsumer(consumerID string, handle func(data *protocol.Any) error) //注册消息消费者
//	//UNRegisterConsumer(consumerID string)  //注销消息消费者
//
//	AsyncBroadcast(serviceType string, message proto.Message) error
//	AsyncPush(serviceType string, serviceID string, message proto.Message) error //异步推送
//
//	SyncRequest(serviceType string, message proto.Message) (interface{}, error)
//	SyncRequestNode(serviceType string, serviceID string, message proto.Message) (interface{}, error)
//	Request(serviceType string, message interface{}) (interface{}, error)
//	RequestNode(serviceType string, serviceID string, message interface{}) (interface{}, error)
//}

func Init() {
	center.ClusterCenter.ConnectCluster(conf.Config.Cluster)
}

func Close() {
	center.ClusterCenter.Close()
}



