/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/4/12
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package dispatch

import (
	"aliens/cluster/message"
	"github.com/gogo/protobuf/proto"
	"aliens/protocol/base"
)


func newGRPCDispatcher() *GRPCDispatcher {
	return &GRPCDispatcher{make(map[string]*message.RemoteService)}
}

type GRPCDispatcher struct {
	serviceMapping map[string]*message.RemoteService
}

//阻塞请求消息 - 根据负载均衡动态分配一个节点处理
func (dispatcher *GRPCDispatcher) SyncRequest(serviceName string, message proto.Message) (interface{}, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	request := &base.Any{Value: data}
	return dispatcher.Request(serviceName, request)
}

//同步阻塞请求
func (dispatcher *GRPCDispatcher) SyncRequestNode(serviceName string, serviceID string, message proto.Message) (interface{}, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	request := &base.Any{Value: data}
	return dispatcher.RequestNode(serviceName, serviceID, request)
}

////同步推送
//func SyncPush(serviceName string, serviceID string, message proto.Message) error {
//	_, err := SyncRequestNode(serviceName, serviceID, message)
//	return err
//}
//
////同步阻塞广播
//func SyncBroadcast(serviceName string, message proto.Message) error {
//	data, err := proto.Marshal(message)
//	if err != nil {
//		return err
//	}
//	request := &protocol.Any{Value: data}
//	service := allocService(serviceName)
//	service.BroadcastAll(request)
//	return nil
//}

func (dispatcher *GRPCDispatcher) Request(serviceName string, message interface{}) (interface{}, error) {
	service := dispatcher.allocService(serviceName)
	return service.HandleMessage(message, "")
}

func (dispatcher *GRPCDispatcher) RequestNode(serviceName string, serviceID string, message interface{}) (interface{}, error) {
	service := dispatcher.allocService(serviceName)
	return service.HandleRemoteMessage(serviceID, message)
}

func (dispatcher *GRPCDispatcher) allocService(serviceName string) *message.RemoteService {
	service := dispatcher.serviceMapping[serviceName]
	if service == nil {
		service = message.NewRemoteService(serviceName)
		dispatcher.serviceMapping[serviceName] = service
	}
	return service
}
