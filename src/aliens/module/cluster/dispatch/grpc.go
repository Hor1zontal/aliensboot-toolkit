/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/8/22
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package dispatch


import (
	"aliens/protocol/base"
	"aliens/cluster/message"
	"aliens/protocol"
	"github.com/gogo/protobuf/proto"
	"aliens/cluster/center/service"
)


var serviceMapping = make(map[string]*message.RemoteService)

//阻塞请求消息 - 根据负载均衡动态分配一个节点处理
func RequestMessage(serviceName string, message  *protocol.Request, hashKey string) (*protocol.Response, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	request := &base.Any{Value: data}
	response, err := Request(serviceName, request, hashKey)
	if err != nil {
		return nil, err
	}
	messageRet := &protocol.Response{}
	messageRet.Unmarshal(response.GetValue())
	return  messageRet, nil
}

//异步发送信息
func SendNodeMessage(serviceName string, serviceID string, message *protocol.Request) error {
	data, _ := message.Marshal()
	request := &base.Any{Value: data}
	return SendNode(serviceName, serviceID, request)
}


//同步阻塞请求
func RequestNodeMessage(serviceName string, serviceID string, message *protocol.Request) (*protocol.Response, error) {
	data, _ := message.Marshal()
	request := &base.Any{Value: data}
	response, err := RequestNode(serviceName, serviceID, request)
	if err != nil {
		return nil, err
	}
	messageRet := &protocol.Response{}
	messageRet.Unmarshal(response.GetValue())
	return  messageRet, nil
}

func Request(serviceName string, message *base.Any, hashKey string) (*base.Any, error) {
	service := allocService(serviceName)
	return service.Request(message, hashKey)
}

func RequestNode(serviceName string, serviceID string, message *base.Any) (*base.Any, error) {
	service := allocService(serviceName)
	return service.RequestNode(serviceID, message)
}

func AsyncRequest(serviceName string, message *base.Any, hashKey string, callback service.Callback) error {
	service := allocService(serviceName)
	return service.AsyncRequest(message, hashKey, callback)
}

func AsyncRequestNode(serviceName string, serviceID string, message *base.Any, callback service.Callback) error {
	service := allocService(serviceName)
	return service.AsyncRequestNode(serviceID, message, callback)
}

func Send(serviceName string, message *base.Any, hashKey string) error {
	service := allocService(serviceName)
	return service.Send(message, hashKey)
}

func SendNode(serviceName string, serviceID string, message *base.Any) error {
	service := allocService(serviceName)
	return service.SendNode(serviceID, message)
}

func allocService(serviceName string) *message.RemoteService {
	service := serviceMapping[serviceName]
	if service == nil {
		service = message.NewRemoteService(serviceName)
		serviceMapping[serviceName] = service
	}
	return service
}





