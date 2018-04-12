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
	"aliens/module/cluster/cache"
	"github.com/pkg/errors"
	"fmt"
	"aliens/module/cluster"
	"github.com/gogo/protobuf/proto"
	"aliens/protocol"
)

//url - service
var serviceMapping = make(map[string]*message.RemoteService)


//网关推送信息
func GatePush(clientID string, messageID uint32, message proto.Message) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	request := &protocol.Any{SessionId:clientID, MessageId: messageID, Value:data}
	return push(clientID, request)
}

func MessageRequest(serviceType string, message proto.Message) (interface{}, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	request := &protocol.Any{Value:data}
	return Request(serviceType, request)
}

func MessageRequestNode(serviceType string, serviceID string, message proto.Message) (interface{}, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	request := &protocol.Any{Value:data}
	return RequestNode(serviceType, serviceID, request)
}

func Request(serviceType string, message interface{}) (interface{}, error) {
	service := allocService(serviceType)
	return service.HandleMessage(message)
}

func RequestNode(serviceType string, serviceID string, message interface{}) (interface{}, error) {
	service := allocService(serviceType)
	return service.HandleRemoteMessage(serviceID, message)
}

//网关推送消息
func push(clientID string, message interface{}) error {
	gateID := cache.ClusterCache.GetClientGateID(clientID)
	if gateID == "" {
		return errors.New(fmt.Sprint("gate ID can not found, clientID : %v", clientID))
	}
	_, err := RequestNode(cluster.SERVICE_GATE, gateID, message)
	return err
}

func allocService(serviceID string) *message.RemoteService {
	service := serviceMapping[serviceID]
	if service == nil {
		service = message.NewRemoteService(serviceID)
		serviceMapping[serviceID] = service
	}
	return service
}
