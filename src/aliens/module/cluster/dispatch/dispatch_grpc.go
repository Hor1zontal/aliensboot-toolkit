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
	"aliens/protocol"
	"errors"
)


func newGRPCDispatcher() *GRPCDispatcher {
	return &GRPCDispatcher{make(map[string]*message.RemoteService)}
}

type GRPCDispatcher struct {
	serviceMapping map[string]*message.RemoteService
}

//阻塞请求消息 - 根据负载均衡动态分配一个节点处理
func (dispatcher *GRPCDispatcher) SyncRequest(serviceName string, message  *protocol.Request, hashKey string) (*protocol.Response, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	request := &base.Any{Value: data}
	response, err := dispatcher.Request(serviceName, request, hashKey)
	if err != nil {
		return nil, err
	}
	any, ok := response.(*base.Any)
	if !ok {
		return nil, errors.New("invalid rpc ret data")
	}
	messageRet := &protocol.Response{}
	messageRet.Unmarshal(any.GetValue())
	return  messageRet, nil
}

//同步阻塞请求
func (dispatcher *GRPCDispatcher) SyncRequestNode(serviceName string, serviceID string, message *protocol.Request) (*protocol.Response, error) {
	data, _ := message.Marshal()
	request := &base.Any{Value: data}
	response, err := dispatcher.RequestNode(serviceName, serviceID, request)
	if err != nil {
		return nil, err
	}
	any, ok := response.(*base.Any)
	if !ok {
		return nil, errors.New("invalid rpc ret data")
	}
	messageRet := &protocol.Response{}
	messageRet.Unmarshal(any.GetValue())
	return  messageRet, nil

}

func (dispatcher *GRPCDispatcher) Request(serviceName string, message interface{}, hashKey string) (interface{}, error) {
	service := dispatcher.allocService(serviceName)
	return service.HandleMessage(message, hashKey)
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
